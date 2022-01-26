yum install perl* -y


https://www.aso.com.tr/blog/233-adsl-kullanicilari-smtp-port-degisimi-587.html

reverse dns kaydı 
https://www.beyaz.net/tr/ipucu/entry/136/reverse-dns-kaydi-olusturmak-ve-kontrol-etmek

komutlar 
sudo -u zimbra -i

zmcontrol status

sudo su - zimbra -c "zmcontrol status"
sudo su - zimbra -c "zmcontrol restart"
sudo su - root 
https://93.115.79.177:7071/

ssh root@93.115.79.177  WZeXDKmy2Tx9FvX

yum  install nano -y

nano /etc/hosts

93.115.79.177 mail.erban.com.tr

sudo hostnamectl set-hostname mail.erban.com.tr  --static



dkim kurulumu 
https://wiki.zimbra.com/wiki/Configuring_for_DKIM_Signing

spf kaydı 
v=spf1 a mx a:mail.erban.com.tr a:erban.com.tr ip4:93.115.79.177 all 

DMARC KAYDI key adı _dmarc
v=DMARC1; p=quarantine; rua=mailto:admin@erban.com.tr; ruf=mailto:admin@erban.com.tr; sp=quarantine

outlok ssl 
https://www.mshowto.org/zimbra-mail-servere-ucretsiz-ssl-kurulumu-lets-encrypt.html

centos 8 
https://computingforgeeks.com/install-zimbra-mail-server-on-centos-rhel/

ssl çözümü
https://www.sbarjatiya.com/notes_wiki/index.php/CentOS_7.x_Install_lets_encrypt_automated_SSL_certificate_in_Zimbra


centos 7 kurulumu 
https://www.unixmen.com/install-zimbra-collaboration-suite-8-6-0-centos-7/


https://lorenzo.mile.si/letsencrypt-zimbra-the-easy-way/242/

certbot_zimbra.sh -n -j  --prompt-confirm

certbot upgrade 
https://www.zone1creative.co.uk/upgrading-certbot-on-centos-7-letsencrypt/


./letsencrypt-auto certonly --standalone -n -j  --prompt-confirm

https://www.zimbra.org/download/zimbra-collaboration
wget https://files.zimbra.com/downloads/8.8.15_GA/zcs-8.8.15_GA_3953.RHEL8_64.20200629025823.tgz


 apachectl -S

The nginx config can be tested with the following:

 mkdir -p /opt/zimbra/data/nginx/html/.well-known/acme-challenge
 echo "test" > /opt/zimbra/data/nginx/html/.well-known/acme-challenge/test.txt

 sudo su restart
 sudo -u zimbra /opt/zimbra/bin/zmproxyctl restart


 zmprov ms mail.erban.com.tr zimbraReverseProxyMailMode redirect


bu çözebilir 
https://syslint.com/blog/tutorial/how-to-install-lets-encrypt-ssl-with-zimbra-fully-automated-configuration/

ssl force 
https://mellowhost.com/blog/how-to-redirect-http-to-https-zimbra-8-8.html





ssl için ders 
https://www.youtube.com/watch?v=qi0yt8TdxE4&ab_channel=NETWORLD

der 1 https://www.youtube.com/watch?v=kNfxANwb6BU&ab_channel=AryyaW 
der2 https://www.youtube.com/watch?v=Ghx3zRVGPUQ&ab_channel=AryyaW

# BUNU DENE 
https://www.youtube.com/watch?v=Ug8S9CsAvpI&ab_channel=AaFikry

# cerbot 8 

Let’s Encrypt Zimbra 8.8.15 CentOS 8

-dnf update
-dnf install wget

Instalando Certbot

-wget https://dl.eff.org/certbot-auto

-chmod +x certbot-auto
-mv certbot-auto /usr/local/bin

Detenemos Zimbra

-su zimbra
-zmproxyctl stop
-exit

export EMAIL="admin@erban.com.tr"
certbot-auto certonly --standalone  -n -j  --prompt-confirm \
-d mail.erban.com.tr \
--preferred-challenges http \
--agree-tos \
-m $EMAIL \
--keep-until-expiring

Ver el certificado:

-ls -lh /etc/letsencrypt/live/mail.erban.com.tr

Crear el directorio para el certificado:

-mkdir /opt/zimbra/ssl/letsencrypt
-cp /etc/letsencrypt/live/mta.ecim.co.cu/* /opt/zimbra/ssl/letsencrypt/

Vemos que todo este en orden:
-ls /opt/zimbra/ssl/letsencrypt/

Copiamos el contenido de:
https://letsencrypt.org/certs/trustid-x3-root.pem.txt

Editamos:

-nano /opt/zimbra/ssl/letsencrypt/chain.pem

Y pegamos debajo el contenido de https://letsencrypt.org/certs/trustid-x3-root.pem.txt

Establecemos permisos:

-chown -R zimbra:zimbra /opt/zimbra/ssl/letsencrypt/

Verificamos:

-ls -lha /opt/zimbra/ssl/letsencrypt/

Entramos como usuario Zimbra:

-su zimbra

Verificamos que todo este correcto:

-/opt/zimbra/bin/zmcertmgr verifycrt comm /opt/zimbra/ssl/letsencrypt/privkey.pem /opt/zimbra/ssl/letsencrypt/cert.pem /opt/zimbra/ssl/letsencrypt/chain.pem

Salimos de usuario Zimbra

-exit

Hacemos salva del certificado actual, principalmente si vamos a renovar antes de que expire el anterior

-cp -a /opt/zimbra/ssl/zimbra /opt/zimbra/ssl/zimbra.$(date "+%Y%.m%.d-%H.%M")

Copiamos al directorio de certificados de zimbra y damos permiso:

-cp /opt/zimbra/ssl/letsencrypt/privkey.pem /opt/zimbra/ssl/zimbra/commercial/commercial.key
-chown zimbra:zimbra /opt/zimbra/ssl/zimbra/commercial/commercial.key
