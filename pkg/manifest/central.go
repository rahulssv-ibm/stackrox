package manifest

import (
	"context"
	"fmt"

	"github.com/stackrox/rox/operator/pkg/types"
	"github.com/stackrox/rox/pkg/certgen"
	"github.com/stackrox/rox/pkg/mtls"
	"github.com/stackrox/rox/pkg/renderer"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (m *manifestGenerator) applyCentral(ctx context.Context) error {
	err := m.createServiceAccount(ctx, "central")
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create central service account: %w\n", err)
	}
	log.Info("Created central service account")

	err = m.createServiceAccount(ctx, "central-db")
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create central-db service account: %w\n", err)
	}
	log.Info("Created central-db service account")

	err = m.createCentralEndpointsConfig(ctx)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create central endpoints config: %w\n", err)
	}
	log.Info("Created central endpoints config")

	err = m.createCentralConfig(ctx)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create central config: %w\n", err)
	}
	log.Info("Created central config")

	err = m.createCentralDbConfig(ctx)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create central db config: %w\n", err)
	}
	log.Info("Created central db config")

	err = m.createAdminPassword(ctx)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create admin password: %w\n", err)
	}

	err = m.createTlsSecrets(ctx)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failed to create TLS secret: %w\n", err)
	}

	err = m.applyCentralDbDeployment(ctx)
	if err != nil {
		return err
	}

	err = m.applyCentralDeployment(ctx)
	if err != nil {
		return err
	}

	err = m.applyCentralServices(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *manifestGenerator) createCentralDbConfig(ctx context.Context) error {
	cm := v1.ConfigMap{
		Data: map[string]string{
			"pg_hba.conf": `local   all             all                                     scram-sha-256
host    all             all             127.0.0.1/32            scram-sha-256
host    all             all             ::1/128                 scram-sha-256
local   replication     all                                     trust
host    replication     all             127.0.0.1/32            trust
host    replication     all             ::1/128                 trust

hostnossl  all       all   0.0.0.0/0     reject
hostnossl  all       all   ::0/0         reject

hostssl    all       all   0.0.0.0/0     scram-sha-256
hostssl    all       all   ::0/0         scram-sha-256`,
			"postgresql.conf": `hba_file = '/etc/stackrox.d/config/pg_hba.conf'
listen_addresses = '*'
max_connections = 200
password_encryption = scram-sha-256

ssl = on
ssl_ca_file = '/run/secrets/stackrox.io/certs/root.crt'
ssl_cert_file = '/run/secrets/stackrox.io/certs/server.crt'
ssl_key_file = '/run/secrets/stackrox.io/certs/server.key'

shared_buffers = 2GB
work_mem = 40MB
maintenance_work_mem = 512MB
effective_cache_size = 4GB

dynamic_shared_memory_type = posix
max_wal_size = 5GB
min_wal_size = 80MB

log_timezone = 'Etc/UTC'
datestyle = 'iso, mdy'
timezone = 'Etc/UTC'
lc_messages = 'en_US.utf8'
lc_monetary = 'en_US.utf8'
lc_numeric = 'en_US.utf8'
lc_time = 'en_US.utf8'

default_text_search_config = 'pg_catalog.english'
shared_preload_libraries = 'pg_stat_statements'`,
		},
	}
	cm.SetName("central-db-config")
	_, err := m.Client.CoreV1().ConfigMaps(m.Namespace).Create(ctx, &cm, metav1.CreateOptions{})

	return err
}

func (m *manifestGenerator) createCentralConfig(ctx context.Context) error {
	cm := v1.ConfigMap{
		Data: map[string]string{
			"central-config.yaml": `maintenance:
  safeMode: false # When set to true, Central will sleep forever on the next restart
  compaction:
    enabled: true
    bucketFillFraction: .5 # This controls how densely to compact the buckets. Usually not advised to modify
    freeFractionThreshold: 0.75 # This is the threshold for free bytes / total bytes after which compaction will occur
  forceRollbackVersion: none # This is the config and target rollback version after upgrade complete.`,
		},
	}
	cm.SetName("central-config")
	_, err := m.Client.CoreV1().ConfigMaps(m.Namespace).Create(ctx, &cm, metav1.CreateOptions{})

	cm = v1.ConfigMap{
		Data: map[string]string{
			"central-external-db.yaml": `centralDB:
   external: false
   source: >
     host=central-db.stackrox.svc
     port=5432
     user=postgres
     sslmode=verify-ca
     sslrootcert=/run/secrets/stackrox.io/certs/ca.pem
     statement_timeout=1.2e+06
     pool_min_conns=10
     pool_max_conns=90
     client_encoding=UTF8`,
		},
	}
	cm.SetName("central-external-db")
	_, err = m.Client.CoreV1().ConfigMaps(m.Namespace).Create(ctx, &cm, metav1.CreateOptions{})

	return err
}

func (m *manifestGenerator) createCentralEndpointsConfig(ctx context.Context) error {
	cm := v1.ConfigMap{
		Data: map[string]string{
			"endpoints.yaml": "",
		},
	}
	cm.SetName("central-endpoints")
	_, err := m.Client.CoreV1().ConfigMaps(m.Namespace).Create(ctx, &cm, metav1.CreateOptions{})

	return err
}

func (m *manifestGenerator) createTlsSecrets(ctx context.Context) error {
	err := m.applyTlsSecret(ctx, "central-tls", func(fileMap types.SecretDataMap) error {
		if err := certgen.IssueCentralCert(fileMap, m.CA, mtls.WithNamespace(m.Namespace)); err != nil {
			return fmt.Errorf("issuing central service certificate: %w\n", err)
		}

		jwtKey, err := certgen.GenerateJWTSigningKey()
		if err != nil {
			return fmt.Errorf("generating JWT signing key: %w\n", err)
		}

		certgen.AddJWTSigningKeyToFileMap(fileMap, jwtKey)
		return nil
	})

	if err != nil {
		return err
	}

	return m.applyTlsSecret(ctx, "central-db-tls", func(fileMap types.SecretDataMap) error {
		subjects := []mtls.Subject{mtls.CentralDBSubject}
		if err := certgen.IssueOtherServiceCerts(fileMap, m.CA, subjects, mtls.WithNamespace(m.Namespace)); err != nil {
			return fmt.Errorf("issuing central service certificate: %w\n", err)
		}

		return nil
	})
}

// TODO: Use this in one of the options
func (m *manifestGenerator) createCentralDbPvc(ctx context.Context) error {
	pvc := v1.PersistentVolumeClaim{
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{"ReadWriteOnce"},
			Resources: v1.VolumeResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse("1Gi"),
				},
			},
		},
	}
	pvc.SetName("central-db")
	_, err := m.Client.CoreV1().PersistentVolumeClaims(m.Namespace).Create(ctx, &pvc, metav1.CreateOptions{})

	return err
}

func (m *manifestGenerator) createAdminPassword(ctx context.Context) error {
	password := "letmein"

	applyPassword := func(fileMap types.SecretDataMap) error {
		fileMap["password"] = []byte(password)
		return nil
	}

	if err := m.applyTlsSecret(ctx, "admin-pass", applyPassword); err != nil {
		return err
	}

	if err := m.applyTlsSecret(ctx, "central-db-password", applyPassword); err != nil {
		return err
	}

	return m.applyTlsSecret(ctx, "central-htpasswd", func(fileMap types.SecretDataMap) error {
		htpasswdBytes, err := renderer.CreateHtpasswd(password)
		if err != nil {
			return err
		}

		fileMap["htpasswd"] = htpasswdBytes
		return nil
	})
}

func (m *manifestGenerator) applyCentralDbDeployment(ctx context.Context) error {
	deployment := apps.Deployment{
		Spec: apps.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "central-db",
				},
			},
			Strategy: apps.DeploymentStrategy{
				Type: apps.RecreateDeploymentStrategyType,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "central-db",
					},
				},
				Spec: v1.PodSpec{
					ServiceAccountName: "central-db",
					SecurityContext: &v1.PodSecurityContext{
						FSGroup: &PostgresUser,
						SeccompProfile: &v1.SeccompProfile{
							Type: v1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []v1.Container{{
						Name:            "central-db",
						Image:           "quay.io/stackrox-io/central-db:latest",
						SecurityContext: RestrictedSecurityContext(PostgresUser),
						Ports: []v1.ContainerPort{{
							Name:          "postgresql",
							ContainerPort: 5432,
							Protocol:      v1.ProtocolTCP,
						}},
						Env: []v1.EnvVar{
							{
								Name:  "POSTGRES_HOST_AUTH_METHOD",
								Value: "password",
							},
							{
								Name:  "PGDATA",
								Value: "/var/lib/postgresql/data/pgdata",
							},
						},
					}},
					InitContainers: []v1.Container{{
						Name:            "init-db",
						Image:           "quay.io/stackrox-io/central-db:latest",
						Command:         []string{"init-entrypoint.sh"},
						SecurityContext: RestrictedSecurityContext(PostgresUser),
						Env: []v1.EnvVar{
							{
								Name:  "PGDATA",
								Value: "/var/lib/postgresql/data/pgdata",
							},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:      "disk",
								MountPath: "/var/lib/postgresql/data",
							},
						},
					}},
				},
			},
		},
	}
	volumeMounts := []VolumeDefAndMount{
		{
			Name:      "config-volume",
			MountPath: "/etc/stackrox.d/config/",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: "central-db-config",
						},
					},
				},
			},
		},
		{
			Name:      "disk",
			MountPath: "/var/lib/postgresql/data",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
				// VolumeSource: v1.VolumeSource{
				// 	PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
				// 		ClaimName: "central-db",
				// 	},
				// },
			},
		},
		{
			Name:      "central-db-tls-volume",
			MountPath: "/run/secrets/stackrox.io/certs",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						DefaultMode: &ReadOnlyMode,
						SecretName:  "central-db-tls",
						Items: []v1.KeyToPath{
							{
								Key:  "central-db-cert.pem",
								Path: "server.crt",
							},
							{
								Key:  "central-db-key.pem",
								Path: "server.key",
							},
							{
								Key:  "ca.pem",
								Path: "root.crt",
							},
						},
					},
				},
			},
		},
		{
			Name:      "shared-memory",
			MountPath: "/dev/shm",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{
						Medium:    v1.StorageMediumMemory,
						SizeLimit: &TwoGigs,
					},
				},
			},
		},
	}

	dbPasswd := VolumeDefAndMount{
		Name:      "central-db-password",
		MountPath: "/run/secrets/stackrox.io/secrets",
		Volume: v1.Volume{
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: "central-db-password",
				},
			},
		},
	}

	dbPasswd.Apply(&deployment.Spec.Template.Spec.InitContainers[0], &deployment.Spec.Template.Spec)

	for _, v := range volumeMounts {
		v.Apply(&deployment.Spec.Template.Spec.Containers[0], &deployment.Spec.Template.Spec)
	}

	deployment.SetName("central-db")
	_, err := m.Client.AppsV1().Deployments(m.Namespace).Create(ctx, &deployment, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		_, err = m.Client.AppsV1().Deployments(m.Namespace).Update(ctx, &deployment, metav1.UpdateOptions{})
		log.Info("Updated central-db deployment")
	} else {
		log.Info("Created central-db deployment")
	}

	return err
}

func (m *manifestGenerator) applyCentralDeployment(ctx context.Context) error {
	deployment := apps.Deployment{
		Spec: apps.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "central",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "central",
					},
				},
				Spec: v1.PodSpec{
					ServiceAccountName: "central",
					Containers: []v1.Container{{
						Name: "central",
						// Image:           "quay.io/klape/stackrox:latest",
						Image:           "localhost:5001/stackrox/main:latest",
						ImagePullPolicy: v1.PullAlways,
						Command:         []string{"/stackrox/central-entrypoint.sh"},
						Ports: []v1.ContainerPort{{
							Name:          "api",
							ContainerPort: 8443,
							Protocol:      v1.ProtocolTCP,
						}},
						Env: []v1.EnvVar{
							{
								Name: "ROX_NAMESPACE",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										FieldPath: "metadata.namespace",
									},
								},
							},
						},
					}},
				},
			},
		},
	}

	trueBool := true
	volumeMounts := []VolumeDefAndMount{
		{
			Name:      "varlog",
			MountPath: "/var/log/stackrox/",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
			},
		},
		{
			Name:      "central-tmp-volume",
			MountPath: "/tmp",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
			},
		},
		{
			Name:      "central-etc-ssl-volume",
			MountPath: "/etc/ssl",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
			},
		},
		{
			Name:      "central-etc-pki-volume",
			MountPath: "/etc/pki/ca-trust",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
			},
		},
		{
			Name:      "central-certs-volume",
			MountPath: "/run/secrets/stackrox.io/certs/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						DefaultMode: &ReadOnlyMode,
						SecretName:  "central-tls",
					},
				},
			},
		},
		{
			Name:      "central-default-tls-cert-volume",
			MountPath: "/run/secrets/stackrox.io/default-tls-cert/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "central-default-tls-cert",
						Optional:   &trueBool,
					},
				},
			},
		},
		{
			Name:      "central-htpasswd-volume",
			MountPath: "/run/secrets/stackrox.io/htpasswd/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "central-htpasswd",
						Optional:   &trueBool,
					},
				},
			},
		},
		{
			Name:      "central-jwt-volume",
			MountPath: "/run/secrets/stackrox.io/jwt/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "central-tls",
						Items: []v1.KeyToPath{
							{
								Key:  "jwt-key.pem",
								Path: "jwt-key.pem",
							},
						},
					},
				},
			},
		},
		{
			Name:      "additional-ca-volume",
			MountPath: "/usr/local/share/ca-certificates/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "additional-ca",
						Optional:   &trueBool,
					},
				},
			},
		},
		{
			Name:      "central-license-volume",
			MountPath: "/run/secrets/stackrox.io/central-license/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "central-license",
						Optional:   &trueBool,
					},
				},
			},
		},
		{
			Name:      "central-config-volume",
			MountPath: "/etc/stackrox",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: "central-config",
						},
					},
				},
			},
		},
		{
			Name:      "proxy-config-volume",
			MountPath: "/run/secrets/stackrox.io/proxy-config/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "proxy-config",
						Optional:   &trueBool,
					},
				},
			},
		},
		{
			Name:      "endpoints-config-volume",
			MountPath: "/etc/stackrox.d/endpoints/",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: "central-endpoints",
						},
					},
				},
			},
		},
		{
			Name:      "central-db-password",
			MountPath: "/run/secrets/stackrox.io/db-password",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "central-db-password",
					},
				},
			},
		},
		{
			Name:      "stackrox-db",
			MountPath: "/var/lib/stackrox",
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{},
				},
			},
		},
		{
			Name:      "central-external-db-volume",
			MountPath: "/etc/ext-db",
			ReadOnly:  true,
			Volume: v1.Volume{
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						DefaultMode: &ReadOnlyMode,
						LocalObjectReference: v1.LocalObjectReference{
							Name: "central-external-db",
						},
					},
				},
			},
		},
	}

	for _, v := range volumeMounts {
		v.Apply(&deployment.Spec.Template.Spec.Containers[0], &deployment.Spec.Template.Spec)
	}

	deployment.SetName("central")

	_, err := m.Client.AppsV1().Deployments(m.Namespace).Create(ctx, &deployment, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		_, err = m.Client.AppsV1().Deployments(m.Namespace).Update(ctx, &deployment, metav1.UpdateOptions{})
		log.Info("Updated central deployment")
	} else {
		log.Info("Created central deployment")
	}

	return err
}

func (m *manifestGenerator) applyCentralServices(ctx context.Context) error {
	// central

	svc := v1.Service{
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "central",
			},
			Ports: []v1.ServicePort{{
				Name:       "https",
				Port:       8443,
				Protocol:   v1.ProtocolTCP,
				TargetPort: intstr.FromString("api"),
			}},
		},
	}

	svc.SetName("central")

	_, err := m.Client.CoreV1().Services(m.Namespace).Create(ctx, &svc, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		_, err = m.Client.CoreV1().Services(m.Namespace).Update(ctx, &svc, metav1.UpdateOptions{})
		log.Info("Updated central service")
	} else {
		log.Info("Created central service")
	}

	if err != nil {
		return err
	}

	// central-db

	svc = v1.Service{
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "central-db",
			},
			Ports: []v1.ServicePort{{
				Name:       "tcp-db",
				Port:       5432,
				Protocol:   v1.ProtocolTCP,
				TargetPort: intstr.FromString("postgresql"),
			}},
		},
	}

	svc.SetName("central-db")

	_, err = m.Client.CoreV1().Services(m.Namespace).Create(ctx, &svc, metav1.CreateOptions{})

	if errors.IsAlreadyExists(err) {
		_, err = m.Client.CoreV1().Services(m.Namespace).Update(ctx, &svc, metav1.UpdateOptions{})
		log.Info("Updated central-db service")
	} else {
		log.Info("Created central-db service")
	}

	return err
}
