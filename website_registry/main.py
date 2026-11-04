import re
import socket

def get_ip(domain):
    try:
        return socket.gethostbyname(domain)
    except socket.gaierror:
        return None


def extract_domain(url):
    m = re.match(r"^(?:https?:\/\/)?(?:www\.)?([^\/:?#]+)", url)
    return m.group(1) if m else None


def registry_url(url):
    domain = extract_domain(url)
    if not domain:
        print("Domain didn't get")
        return

    with open("register.txt", "r", encoding="utf-8") as file:
        for line in file:
            parts = line.rstrip("\n").split(";")
            urls    = [u.strip() for u in parts[1].split(",") if u.strip()]
            domains = [d.strip() for d in parts[2].split(",") if d.strip()]
            if url in urls or domain in domains:
                print("block")
                return

    ip = get_ip(domain)
    if ip:
        with open("register.txt", "r", encoding="utf-8") as file:
            for line in file:
                parts = line.rstrip("\n").split(";")
                ips = [i.strip() for i in parts[3].split(",") if i.strip()]
                if ip in ips:
                    print("block")
                    return

    print("unlocked")


url = input("Enter urls: ")
registry_url(url)