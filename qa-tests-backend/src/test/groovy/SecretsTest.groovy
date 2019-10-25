import services.SecretService
import org.junit.experimental.categories.Category
import groups.BAT
import objects.Deployment
import io.stackrox.proto.storage.SecretOuterClass.Secret
import util.Timer

class SecretsTest extends BaseSpecification {

    private static Deployment renderDeployment(String deploymentName, String secretName) {
        return new Deployment()
                .setName (deploymentName)
                .setNamespace("qa")
                .setImage ("busybox")
                .addLabel ( "app", "test" )
                .addVolume("test", "/etc/try")
                .addSecretName("test", secretName)
                .setCommand(["sleep", "600"])
    }

    @Category(BAT)
    def "Verify the secret api can return the secret's information when adding a new secret"() {
        when:
        "Create a Secret"
        String secretName = "qasec"
        String secID = orchestrator.createSecret(secretName)

        then:
        "Verify Secret is added to the list"
        assert SecretService.getSecret(secID) != null

        cleanup:
        "Remove Secret #secretName"
        orchestrator.deleteSecret(secretName)
    }

    @Category(BAT)
    def "Verify the secret item should show the binding deployments"() {
        when:
        "Create a Secret"
        String secretName = "qasec"
        String secID = orchestrator.createSecret("qasec")

        and:
        "Create a Deployment using above created secret"
        String deploymentName = "depwithsecrets"
        Deployment deployment = renderDeployment(deploymentName, secretName)
        orchestrator.createDeployment(deployment)

        then:
        "Verify the deployment is binding with the secret"
        assert SecretService.getSecret(secID) != null
        Set<String> secretSet = orchestrator.getDeploymentSecrets(deployment)
        assert secretSet.contains(secretName)

        cleanup:
        "Remove Secret #secretName and Deployment #deploymentName"
        orchestrator.deleteAndWaitForDeploymentDeletion(deployment)
        orchestrator.deleteSecret(secretName)
    }

    @Category(BAT)
    def "Verify the secret should not show the deleted binding deployment"() {
        when:
        "Create a Secret and bind deployment with it"
        String secretName = "qasec"
        String deploymentName = "depwithsecrets"
        String secID = orchestrator.createSecret("qasec")
        Deployment deployment = renderDeployment(deploymentName, secretName)

        orchestrator.createDeployment(deployment)

        Secret secretInfo = SecretService.getSecret(secID)
        int preNum = secretInfo.getRelationship().getDeploymentRelationshipsCount()

        and:
        "Delete the binding deployment"
        orchestrator.deleteAndWaitForDeploymentDeletion(deployment)

        then:
        "Verify the binding deployment is gone from the secret"
        Secret secretUpdate = null
        def timer = new Timer(10, 3)

        //Add waiting logic cause stackrox need some time to response the number of deployments' change
        while (timer.IsValid()) {
            secretUpdate = SecretService.getSecret(secID)
            if (secretUpdate.getRelationship().getDeploymentRelationshipsCount() == (preNum - 1)) {
                break
            }
        }

        assert secretUpdate.getRelationship().getDeploymentRelationshipsCount() == (preNum - 1)

        cleanup:
        "Remove Secret #secretName"
        orchestrator.deleteSecret(secretName)
    }

    @Category(BAT)
    def "Verify the secret information should not be infected by the previous secrets"() {
        when:
        "Create a Secret and bind deployment with it"
        String secretName = "qasec"
        String deploymentName = "depwithsecrets"

        String secID = orchestrator.createSecret("qasec")
        Deployment deployment = renderDeployment(deploymentName, secretName)
        orchestrator.createDeployment(deployment)

        and:
        "Delete this deployment and create another deployment binding with the secret name with different name"
        orchestrator.deleteAndWaitForDeploymentDeletion(deployment)

        String deploymentSecName = "depwithsecretssec"
        Deployment deploymentSec = renderDeployment(deploymentSecName, secretName)
        orchestrator.createDeployment(deploymentSec)

        then:
        "Verify the secret should show the new bounding deployment"
        Secret secretInfo = SecretService.getSecret(secID)
        assert secretInfo.getRelationship().getDeploymentRelationshipsCount() == 1
        assert secretInfo.getRelationship().getDeploymentRelationships(0).getName() == deploymentSecName

        cleanup:
        "Remove Deployment #deploymentName and Secret #secretName"
        orchestrator.deleteAndWaitForDeploymentDeletion(deploymentSec)
        orchestrator.deleteSecret(secretName)
    }

    @Category(BAT)
    def "Verify secrets page should not be messed up when a deployment's secret changed"() {
        when:
        "Create a Secret and bind deployment with it"
        String secretNameOne = "qasec1"
        String deploymentNameOne = "depwithsecrets1"
        String secIDOne = orchestrator.createSecret("qasec1")
        Deployment deploymentOne = renderDeployment(deploymentNameOne, secretNameOne)
        orchestrator.createDeployment(deploymentOne)

        String secretNameTwo = "qasec2"
        String deploymentNameTwo = "depwithsecrets2"
        String secIDTwo = orchestrator.createSecret("qasec2")
        Deployment deploymentTwo = renderDeployment(deploymentNameTwo, secretNameTwo)
        orchestrator.createDeployment(deploymentTwo)

        and:
        "Delete this deployment and create another deployment binding with the secret name with different name"
        orchestrator.deleteAndWaitForDeploymentDeletion(deploymentOne, deploymentTwo)

        deploymentOne = renderDeployment(deploymentNameOne, secretNameTwo)
        deploymentTwo = renderDeployment(deploymentNameTwo, secretNameOne)
        orchestrator.createDeployment(deploymentOne)
        orchestrator.createDeployment(deploymentTwo)

        then:
        "Verify the secret should show the new bounding deployment"
        Secret secretInfoOne = SecretService.getSecret(secIDOne)
        Secret secretInfoTwo = SecretService.getSecret(secIDTwo)

        assert secretInfoOne.getRelationship().getDeploymentRelationships(0).getName() == deploymentNameTwo
        assert secretInfoTwo.getRelationship().getDeploymentRelationships(0).getName() == deploymentNameOne

        cleanup:
        "Remove Deployment and Secret"
        orchestrator.deleteAndWaitForDeploymentDeletion(deploymentOne, deploymentTwo)
        orchestrator.deleteSecret(secretNameOne)
        orchestrator.deleteSecret(secretNameTwo)
    }
}
