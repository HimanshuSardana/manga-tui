import requests
h = {
                "Host":                      "chapmanganato.to",
        "User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
        "Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-Language":           "en-US,en;q=0.5",
        "Connection":                "keep-alive",
        "Referer":                   "https://chapmanganato.to/manga-aa951409",
        "Cookie":                    "ci_session=Uj%2FMmsw3snmXDdm%2FAxXvAh30dOWXFaZawBkdrdGCA0eWDAxwI77%2FgH%2BU6TRe4RzUikBDsDMIYQTSDN%2FZ8O388NJzeHTdOsCLpCa6MoysPwk0g9fI1ntRO8qn0%2B3zZHWg6%2Be1SrOYgs0KxZU6wS9lo%2F3dej81aq1Vw%2Baz7EBeSsrYnVVqdcATFl7PhnVh65J3QEvJa8bMKCkeXsdyuGJNCOkGpkkCXlemCTNguS%2F71i2qygsuZa5G4XJqTaSBDWN4%2FzhdrcgGhggfc5wtQIsYT1qYEsXcWXcq2J32x%2BnHMiTp%2FSc9rxbm4jvPnaf8tZOG%2FVBcvbOk3ZvrjxVQiG6bY0V6uYmqM4Fm6eWaj5%2Fhik9dwWz5BJISf%2B4lJJJadzg8CcVIJht5ABZAKrytGkFpKtFhHiXFnKqW4HfV%2Fqt9HoY%3Da3468db3c15b1d69fbb668ff5a3f14ce1c5bffc6; panel-fb-comment=fb-comment-title-show",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest":            "document",
        "Sec-Fetch-Mode":            "navigate",
        "Sec-Fetch-Site":            "same-origin",
        "Sec-Fetch-User":            "?1",
        "Priority":                  "u=1",
        "TE":                        "trailers"}

h2 = {
"Host": "v12.mkklcdnv6tempv4.com",
"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
"Accept": "image/avif,image/webp,*/*",
"Accept-Language": "en-US,en;q=0.5",
"Accept-Encoding": "gzip, deflate, br, zstd",
"Connection": "keep-alive",
"Referer": "https://chapmanganato.to/",
"Sec-Fetch-Dest": "image",
"Sec-Fetch-Mode": "no-cors",
"Sec-Fetch-Site": "cross-site",
"Priority": "u=4"
    }
with open("output.txt", "r") as f:
    ctr = 1
    for link in f:
        if ctr == 1:
            with open("1.webp", "wb") as f2:
                f2.write(requests.get(link, headers=h2).content)
            ctr += 1
        else:
            with open(f"{ctr}.jpg", "wb") as f2:
                f2.write(requests.get(link, headers=h2).content)
            ctr += 1
        print(f"[+] {ctr}.jpg written")

