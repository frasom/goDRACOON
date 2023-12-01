# AppID 

Firstly you will need some more development tools installed for mobile packaging to complete. 
[Fyne Mobile Packaging](https://developer.fyne.io/started/mobile)

## iOS mobile packaging
https://support.staffbase.com/hc/de/articles/115003535352-iOS-App-ID-erstellen

To build iOS apps you will need Xcode installed on your macOS computer as well as the command line tools optional package.

% ~/go/bin/fyne package -os ios -appID com.github.goDRACOON -icon ./resources/drCheck.png
-os=ios requires XCode

### install Xcode Command Line Tools on a Mac:

    * install the full Xcode package
    * install Xcode Command Line Tools when triggered by a command
    * install Xcode Command Line Tools as part of a Homebrew installation.

```
xcode-select --install 
``` 
https://www.freecodecamp.org/news/install-xcode-command-line-tools/


## Android mobile packaging

For Android builds you must have the Android SDK and NDK installed with appropriate environment set up so that the tools (such as adb) can be found on the command line. 

% ~/go/bin/fyne package -os android -appID com.github.goDRACOON -icon ./resources/drCheck.png
no Android NDK found in $ANDROID_HOME/ndk-bundle nor in $ANDROID_NDK_HOME
