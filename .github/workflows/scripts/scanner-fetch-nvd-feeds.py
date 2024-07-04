import json
import os
import argparse
from datetime import datetime
import requests
import gzip
import shutil

def dirpath(s):
    if os.path.isdir(s):
        return s
    raise ValueError(f"{s} is not a valid directory path")

def download_and_save_json(url, year, save_dir):
    try:
        response = requests.get(url, stream=True)
        response.raise_for_status()

        gz_file_path = os.path.join(save_dir, f"nvdcve-1.1-{year}.json.gz")
        json_file_path = os.path.join(save_dir, f"nvdcve-1.1-{year}.json")

        # Save the gzipped file
        with open(gz_file_path, 'wb') as gz_file:
            for chunk in response.iter_content(chunk_size=8192):
                gz_file.write(chunk)

        # Decompress and save the JSON file
        with gzip.open(gz_file_path, 'rb') as f_in:
            with open(json_file_path, 'wb') as f_out:
                shutil.copyfileobj(f_in, f_out)

        # Optionally, remove the gzipped file after decompressing
        os.remove(gz_file_path)
        print(f"Saved {json_file_path}")

    except Exception as e:
        print(f"Error processing {url} for year {year}: {e}")
        if os.path.exists(gz_file_path):
            os.remove(gz_file_path)
        if os.path.exists(json_file_path):
            os.remove(json_file_path)

def transform_json(input_file, output_file):
    with open(input_file, 'r') as f:
        data = json.load(f)

    transformed_items = []

    for item in data['CVE_Items']:
        impact_data = item.get('impact', {})
        base_metric_v2 = impact_data.get('baseMetricV2', {})
        base_metric_v3 = impact_data.get('baseMetricV3', {})

        metrics = {}

        if 'cvssV2' in base_metric_v2:
            cvss_v2_data = base_metric_v2['cvssV2']
            metrics["cvssMetricV2"] = [
                {
                    "source": "nvd@nist.gov",
                    "type": "Primary",
                    "cvssData": {
                        "version": "2.0",
                        "vectorString": cvss_v2_data.get('vectorString', ''),
                        "accessVector": cvss_v2_data.get('accessVector', ''),
                        "accessComplexity": cvss_v2_data.get('accessComplexity', ''),
                        "authentication": cvss_v2_data.get('authentication', ''),
                        "confidentialityImpact": cvss_v2_data.get('confidentialityImpact', ''),
                        "integrityImpact": cvss_v2_data.get('integrityImpact', ''),
                        "availabilityImpact": cvss_v2_data.get('availabilityImpact', ''),
                        "baseScore": cvss_v2_data.get('baseScore', 0)
                    },
                    "baseSeverity": base_metric_v2.get('severity', ''),
                    "exploitabilityScore": base_metric_v2.get('exploitabilityScore', 0),
                    "impactScore": base_metric_v2.get('impactScore', 0),
                    "acInsufInfo": base_metric_v2.get('acInsufInfo', False),
                    "obtainAllPrivilege": base_metric_v2.get('obtainAllPrivilege', False),
                    "obtainUserPrivilege": base_metric_v2.get('obtainUserPrivilege', False),
                    "obtainOtherPrivilege": base_metric_v2.get('obtainOtherPrivilege', False),
                    "userInteractionRequired": base_metric_v2.get('userInteractionRequired', False)
                }
            ]

        if 'cvssV3' in base_metric_v3:
            cvss_v3_data = base_metric_v3['cvssV3']
            metrics["cvssMetricV31"] = [
                {
                    "source": "nvd@nist.gov",
                    "type": "Primary",
                    "cvssData": {
                        "version": "3.1",
                        "vectorString": cvss_v3_data.get('vectorString', ''),
                        "attackVector": cvss_v3_data.get('attackVector', ''),
                        "attackComplexity": cvss_v3_data.get('attackComplexity', ''),
                        "privilegesRequired": cvss_v3_data.get('privilegesRequired', ''),
                        "userInteraction": cvss_v3_data.get('userInteraction', ''),
                        "scope": cvss_v3_data.get('scope', ''),
                        "confidentialityImpact": cvss_v3_data.get('confidentialityImpact', ''),
                        "integrityImpact": cvss_v3_data.get('integrityImpact', ''),
                        "availabilityImpact": cvss_v3_data.get('availabilityImpact', ''),
                        "baseScore": cvss_v3_data.get('baseScore', 0),
                        "baseSeverity": cvss_v3_data.get('baseSeverity', '')
                    },
                    "exploitabilityScore": base_metric_v3.get('exploitabilityScore', 0),
                    "impactScore": base_metric_v3.get('impactScore', 0)
                }
            ]

        if metrics:
            transformed_item = {
                "cve": {
                    "id": item['cve']['CVE_data_meta']['ID'],
                    "published": item['publishedDate'],
                    "lastModified": item['lastModifiedDate'],
                    "descriptions": [{"lang": desc['lang'], "value": desc['value']} for desc in item['cve']['description']['description_data']],
                    "metrics": metrics
                }
            }
            transformed_items.append(transformed_item)

    with open(output_file, 'w') as f_out:
        for transformed_item in transformed_items:
            json.dump(transformed_item, f_out)
            f_out.write("\n")
        print(f"Saved {output_file}")

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        'dirpath',
        help="Path to directory where the NVD data will be saved.",
        type=dirpath
    )
    args = parser.parse_args()
    save_dir = args.dirpath

    base_url = "https://nvd.nist.gov/feeds/json/cve/1.1/"
    first_year = 2002
    current_year = datetime.now().year

    for year in range(first_year, current_year + 1):
        url = f"{base_url}nvdcve-1.1-{year}.json.gz"
        download_and_save_json(url, year, save_dir)
        json_file_path = os.path.join(save_dir, f"nvdcve-1.1-{year}.json")
        transformed_file_path = os.path.join(save_dir, f"{year}.nvd.json")
        transform_json(json_file_path, transformed_file_path)

if __name__ == "__main__":
    main()
