import os, sys
import requests
import multiprocessing

intro = """
                    IG: @godacity_
            DC: https://discord.gg/4jy8khC
                        I <3 U
                        
    Make sure your proxy file is in the IP:PORT scheme.
"""
outro = """
    \n\n\nCheck is completed...........u big gae.
"""

def main(proxy):
    __dir__ = os.getcwd()
    proxy_filtered = open(os.path.join(__dir__, 'proxy_list_filtered.txt'), 'a')
    proxy = {"http": 'http://{}'.format(proxy).strip('\n'), "https": 'https://{}'.format(proxy).strip('\n')}
    
    try:
        ip_raw = requests.get(url, proxies=proxy, timeout=2)
        ip = ip_raw.text
        if ip_raw.status_code == 200:
            print('Public IP:', ip.strip('\n') + ' || Proxy: {}'.format(proxy['http'].replace('http://', '')))
            proxy_filtered.write(proxy['http'].replace('http://', '') + '\n')
    except Exception as e:
        print('{FAILED: ' + proxy['http'].replace('http://', '') + '}')
        pass

print(intro)
url = 'http://ipinfo.io/ip'

if __name__ == "__main__":

    pool = multiprocessing.Pool(processes=20)
    try:
        out = pool.map(main, open(sys.argv[1], 'r').readlines())
    except IndexError:
        print('Specify proxy file [python3 proxy_checker.py proxy_file.txt].')
    except FileNotFoundError:
        print('FILE DOES NOT EXIST.')
    pool.close()
    print(outro)
