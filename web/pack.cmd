rd /S/Q dist
md dist
md dist\_nuxt
md dist\admin
md dist\static
md dist\themes
md dist\xblog
xcopy _nuxt dist\_nuxt /e /y
xcopy admin dist\admin /e /y
xcopy static dist\static /e /y
xcopy themes dist\themes /e /y
xcopy xblog dist\xblog /e /y
copy index.php dist\index.php