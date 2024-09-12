# Maintainer: Develoopeer https://github.com/develoopeer/
pkgname=TrayTranslater
pkgver=0.0.1
pkgrel=1
pkgdesc=""
arch=("x86_64")
url="https://github.com/develoopeer/tray-translater/"
license=('GPL')

build() {
	cd $BUILDDIR/go/
	go build -o ttr
}

package(){
	echo $pkgdir
	cd $BUILDDIR/go/
	install -Dm755 ttr "$pkgdir"/usr/bin/ttr-cli
	# cp ttr $pkgdir
}
