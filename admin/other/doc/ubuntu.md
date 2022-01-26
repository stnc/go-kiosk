rm -rf main.exe
go build main.go
#GOOS=windows GOARCH=amd64 go build main.go

GOOS=linux GOARCH=386 go build main.go

chmod +rwx filename to add permissions.
sudo -u zimbra -i
ubuntu
https://stackoverflow.com/questions/23921117/disable-only-full-group-by
 
 
nano ~/.bash_profile
 
   export GOPATH=$HOME/Projects/emma
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
 
source ~/.bash_profile
 
go env GOPATH
 
https://stackoverflow.com/questions/54456186/how-to-fix-environment-variables-not-working-while-running-from-system-d-service

service için yazılacak olan kısım 
# -------

[Unit]
Description=goweb

[Service]
Type=simple
Restart=always
RestartSec=5s
EnvironmentFile=/home/skurban/public_html/.env
ExecStart=/home/skurban/public_html/main
WorkingDirectory=/home/skurban/public_html/

[Install]
WantedBy=multi-user.target


# --------- 


https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
https://medium.com/@tabvn/deploy-golang-application-on-digital-ocean-server-ubuntu-16-04-b7bf5340ccd9
https://juliensalinas.com/en/develop-deploy-whole-website-golang/



ubuntu open ports yani ççalaışanlar 
sudo netstat -ntlp | grep LISTEN


porta iizin vermek (https://www.e2enetworks.com/help/knowledge-base/how-to-open-ports-on-iptables-in-a-linux-server/)
sudo iptables -A INPUT -p tcp --dport 9990 -j ACCEPT
iptables-save

or -- https://askubuntu.com/questions/119393/how-to-save-rules-of-the-iptables
sudo apt-get install iptables-persistent
sudo netfilter-persistent save
sudo netfilter-persistent reload