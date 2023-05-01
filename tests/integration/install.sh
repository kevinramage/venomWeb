# Prerequisite
sudo su
apt install --assume-yes wget
apt install --assume-yes gpg
apt install --assume-yes unzip
apt install --assume-yes git

# Chrome (/usr/bin/)
wget -qO - https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor -o /usr/share/keyrings/googlechrome-linux-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/googlechrome-linux-keyring.gpg] http://dl.google.com/linux/chrome/deb/ stable main" | tee /etc/apt/sources.list.d/google-chrome.list
apt update
apt install --assume-yes google-chrome-stable
google-chrome -version

# Firefox (/usr/bin/)
apt install --assume-yes firefox-esr
firefox -version

# Edge (/usr/bin/)
wget -qO - https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
install -o root -g root -m 644 microsoft.gpg /etc/apt/trusted.gpg.d/
sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/edge stable main" > /etc/apt/sources.list.d/microsoft-edge-dev.list'
rm microsoft.gpg
apt update
apt install microsoft-edge-stable
microsoft-edge -version

# Opera (/usr/bin/)
wget -qO - https://deb.opera.com/archive.key | gpg --dearmor > opera.gpg
install -o root -g root -m 644 opera.gpg /etc/apt/trusted.gpg.d/
sh -c 'echo "deb https://deb.opera.com/opera-stable/ stable non-free" > /etc/apt/sources.list.d/opera-dev.list'
rm opera.gpg
apt update
apt install --assume-yes opera-stable
opera -version

# Brave (/usr/bin/)
wget -qO - https://brave-browser-apt-release.s3.brave.com/brave-browser-archive-keyring.gpg | gpg --dearmor > brave.gpg
install -o root -g root -m 644 brave.gpg /etc/apt/trusted.gpg.d/
sh -c 'echo "deb https://brave-browser-apt-release.s3.brave.com/ stable main" > /etc/apt/sources.list.d/brave-browser-release.list'
rm brave.gpg
apt update
apt install --assume-yes brave-browser
brave-browser -version

# Download chrome driver 112
wget https://chromedriver.storage.googleapis.com/112.0.5615.49/chromedriver_linux64.zip
unzip chromedriver_linux64.zip

# Download gecko driver 33
wget https://github.com/mozilla/geckodriver/releases/download/v0.33.0/geckodriver-v0.33.0-linux64.tar.gz
tar -xvf geckodriver-v0.33.0-linux64.tar.gz

# Download edge driver 112
wget https://msedgedriver.azureedge.net/112.0.1722.39/edgedriver_linux64.zip
unzip edgedriver_linux64.zip

# Download golang
wget  https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
tar -xvf go1.20.2.linux-amd64.tar.gz  
mv go /usr/local
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
go version

# Download venom
go version
git clone -b kevinramage-web-refacto https://github.com/ovh/venom.git
mv venom venomSource
cd venomSource/cli/venom
go build
mv ./venom ../../..