ADIM ADIM 

burada firewall gibi konuları da anlatmış 
https://www.serdarbayram.net/centos-7-uzerine-zimbra-8-8-12-kurulumu.html

sudo dnf -y update

sudo dnf -y install epel-release dnf-utils
sudo dnf config-manager --enable PowerTools
sudo dnf -y install bash-completion vim curl wget unzip openssh-clients telnet net-tools sysstat perl-core libaio nmap-ncat libstdc++.so.6 bind-utils tar nano

sudo dnf -y install chrony

sudo systemctl enable --now chronyd

sudo chronyc sources

sudo systemctl reboot


nano /etc/hosts

93.115.79.177 mail.erban.com.tr

sudo hostnamectl set-hostname mail.erban.com.tr  --static

wget https://files.zimbra.com/downloads/8.8.15_GA/zcs-8.8.15_GA_3953.RHEL8_64.20200629025823.tgz

tar xfv zcs-8.8.15_GA_3953.RHEL8_64.20200629025823.tgz

cd zcs-8.8.15_GA_3953.RHEL8_64.20200629025823

./install.sh --platform--override








dkim kurulumu 
https://wiki.zimbra.com/wiki/Configuring_for_DKIM_Signing

spf kaydı 
v=spf1 a mx a:mail.erban.com.tr a:erban.com.tr ip4:93.115.79.177 all 

DMARC KAYDI key adı _dmarc
v=DMARC1; p=quarantine; rua=mailto:admin@erban.com.tr; ruf=mailto:admin@erban.com.tr; sp=quarantine


outlok ssl 
https://www.mshowto.org/zimbra-mail-servere-ucretsiz-ssl-kurulumu-lets-encrypt.html


performans için 
http://blog.jeshurun.ca/technology/zimbra-8-7-low-memory-ram-performance-tuning

scrollout 
https://steemit.com/utopian-io/@fightmovies/scrollout-f1-tutorial-or-or-make-antivirus-and-spam-protection

https://www.youtube.com/watch?v=5wwpg4zkjFQ&list=LL5UAeI60E16D9dLuy3CoaOQ&index=1319&ab_channel=MariusGologan&t=35s

zimbra create account 
https://zetcode.com/golang/exec-command/

sudo -u zimbra -i

//admin hesapı açar 
zmprov ca admin1@erban.com.tr 11111  zimbraIsAdminAccount TRUE

export LC_ALL=nb_NO.UTF-8

zmprov ca test5@erban.com.tr password cn "Selman Tunç" displayName "Selman Tunç" givenName "Selman" sn "TUNÇ" zimbraPrefFromDisplay "Selman Tunç" mobile "05452398161" telephoneNumber "Dahili: 1156" homePhone "Kisa Kod: 332" company "Erciyes Teknopark" street "Erciyes Teknopark Yerleşkesi Tekno-3 Binası 2. Kat NO: 28 38039 Melikgazi / KAYSERİ"

zimbra kodları 
https://sites.google.com/site/linuxscooter/linux/mail/zimbra-install

http://docs.zimbra.com/docs/ne/6.0.8/administration_guide/A_app-command-line.20.03.html