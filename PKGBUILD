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
	mkdir -p ~/.config/ttr/
	cd $BUILDDIR/go/
	install -Dm755 ttr "$pkgdir"/usr/bin/ttr-cli
	cd $BUILDDIR/
	cp launch.sh $pkgdir/
}
