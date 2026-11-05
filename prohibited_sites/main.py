import json
import time
import os
import whois
from urllib.parse import urlparse

FILE = "domains"

SKIP_DOMAINS = {
    "localhost", "127.0.0.1", "0.0.0.0", "example.com",
}

SKIP_TLDS = {".local", ".localhost", ".test", ".invalid", ".example"}


def clean_domain(raw: str) -> str:
    if not raw:
        return ""
    d = raw.strip().strip('"').strip("'").strip()
    d = d.lstrip('\\').strip('"').strip()
    if "://" in d:
        try:
            parsed = urlparse(d)
            d = parsed.netloc
        except Exception:
            return ""
    d = d.split("/")[0].split(":")[0].strip().lower()
    if "." not in d:
        return ""
    return d


def fetch_blocked_domains_from_file(path: str):
    if not os.path.exists(path):
        print(f"File didn't find: {path}")
        return []
    print(f"Loaded registry from file: {path}...")
    with open(path, "r", encoding="utf-8") as f:
        data = json.load(f)
    print(f"Successuful loaded {len(data)} recoreds from registry")
    return data



def extract_https_domains(blocked_list):
    domains = set()
    for entry in blocked_list:
        if not isinstance(entry, str):
            continue
        domain = clean_domain(entry)
        if domain:
            domains.add(domain)

    print(f"Found {len(domains)} unique domains.")
    return sorted(domains)


def check_domain_availability(domains):
    available_domains = []
    total = len(domains)
    for i, domain in enumerate(domains, start=1):
        if domain in SKIP_DOMAINS:
            continue
        if any(domain.endswith(tld) for tld in SKIP_TLDS):
            continue

        try:
            w = whois.whois(domain)

            if not w or (not w.registrar and not w.creation_date):
                print(f"[{i}/{total}] FREE: {domain}")
                available_domains.append(domain)
            else:
                pass
        except whois.exceptions.WhoisDomainNotFoundError:
            print(f"[{i}/{total}] FREE (WHOIS error): {domain}")
            available_domains.append(domain)
        except Exception as e:
            pass

        time.sleep(1)

    return available_domains


def main():
    blocked_entries = fetch_blocked_domains_from_file(FILE)
    if not blocked_entries:
        return

    https_domains = extract_https_domains(blocked_entries)
    free_domains = check_domain_availability(https_domains)


if __name__ == "__main__":
    main()