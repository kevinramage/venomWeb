echo "Install chrome browser"
curl https://dl.google.com/dl/chrome/mac/universal/stable/gcem/GoogleChrome.pkg > GoogleChrome.pkg
sudo installer -pkg GoogleChrome.pkg

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

