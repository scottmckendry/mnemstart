# Changelog

## [0.7.3](https://github.com/scottmckendry/mnemstart/compare/v0.7.2...v0.7.3) (2025-01-17)


### Bug Fixes

* **auth:** fix securecookie error when session value is too long ([84b5cda](https://github.com/scottmckendry/mnemstart/commit/84b5cda61cb07007cda3b87b40befcca7752c6f1)), closes [#31](https://github.com/scottmckendry/mnemstart/issues/31)

## [0.7.2](https://github.com/scottmckendry/mnemstart/compare/v0.7.1...v0.7.2) (2024-09-09)


### Bug Fixes

* **ui:** include font-face css in dom ([aa0f543](https://github.com/scottmckendry/mnemstart/commit/aa0f543a3750491fc07d1ec4a6d66956f65b9693))

## [0.7.1](https://github.com/scottmckendry/mnemstart/compare/v0.7.0...v0.7.1) (2024-08-25)


### Bug Fixes

* **keymaps:** custom mappings not immediately available on save ([0b28342](https://github.com/scottmckendry/mnemstart/commit/0b2834286348de62aeebca4008a85582a4e9672f))

## [0.7.0](https://github.com/scottmckendry/mnemstart/compare/v0.6.0...v0.7.0) (2024-08-25)


### Features

* **search:** add search suggestions and selection of engines ([07c2b26](https://github.com/scottmckendry/mnemstart/commit/07c2b2602178e734af1a3c847acf54455c5d6c80))
* **settings:** add setting for toggling search suggestions ([f526619](https://github.com/scottmckendry/mnemstart/commit/f526619fee7b625bd0ed21b4211a38c0573cf2ba))
* **ui:** add help popup and shortcut ([87f31ce](https://github.com/scottmckendry/mnemstart/commit/87f31ce8e54b37096ec5404d5f3ecf1481875976))
* **ui:** style search suggestions ([30c49ed](https://github.com/scottmckendry/mnemstart/commit/30c49ed74790d157724a41a4982349d138dcd910))
* **ui:** update modal styles and layout ([a4d72d8](https://github.com/scottmckendry/mnemstart/commit/a4d72d802bce2b2f36074a9178d90eae943c766f))
* **ui:** update settings modal style ([8b7071a](https://github.com/scottmckendry/mnemstart/commit/8b7071af9d81e0437c0dae9e05c0a23f9552cf9f))

## [0.6.0](https://github.com/scottmckendry/mnemstart/compare/v0.5.0...v0.6.0) (2024-08-24)


### Features

* **auth:** add google & gitlab auth providers ([f30ba9b](https://github.com/scottmckendry/mnemstart/commit/f30ba9bee3260cc314a52eb1aa9a1b28a797f87e))
* **ui:** update login page look and application icons ([994176a](https://github.com/scottmckendry/mnemstart/commit/994176aa8327ce3ac6726769dee8d3dc5307d29d))

## [0.5.0](https://github.com/scottmckendry/mnemstart/compare/v0.4.1...v0.5.0) (2024-08-24)


### Features

* **ui:** add QoL default shortcuts for modals ([5b6cec4](https://github.com/scottmckendry/mnemstart/commit/5b6cec4c0e26c2c1ac3b1044af1d49e4dd238a8e))


### Bug Fixes

* **settings:** add err check for insert query ([6627e9e](https://github.com/scottmckendry/mnemstart/commit/6627e9e42a10cc4fdd6e073447594e22d00421a9))

## [0.4.1](https://github.com/scottmckendry/mnemstart/compare/v0.4.0...v0.4.1) (2024-08-24)


### Bug Fixes

* **settings:** resolve sql bug resulting in duplicated settings ([e191dd5](https://github.com/scottmckendry/mnemstart/commit/e191dd57db2d3a54cfd9866b50239004e8b8d5ad))

## [0.4.0](https://github.com/scottmckendry/mnemstart/compare/v0.3.0...v0.4.0) (2024-08-24)


### Features

* **api:** add recoverer middleware to gracefully handle panics ([9e88d8d](https://github.com/scottmckendry/mnemstart/commit/9e88d8dc1e5799a673e94dc94dec57aa00005014))
* **api:** explicitly set methods for routes ([b6b66c2](https://github.com/scottmckendry/mnemstart/commit/b6b66c2d329af6210c08f6cb38b8ad3743d3437b))
* **auth:** convert auth service to chi middleware ([86597c5](https://github.com/scottmckendry/mnemstart/commit/86597c5e21641f926e60c7bd8902f15795136fd6))
* **cd:** add deployment job to pipeline to update sever image ([a7f981b](https://github.com/scottmckendry/mnemstart/commit/a7f981b58057b1164fb0a2d1e38019c55f20eb47))
* **ci:** publish test results to summary ([e3a00b3](https://github.com/scottmckendry/mnemstart/commit/e3a00b3828a11433af20d9b3fa93f1f6eec90075))
* **keymaps:** show status for incorrectly set leader key ([99f2f96](https://github.com/scottmckendry/mnemstart/commit/99f2f96a1038f5db1b10d3ae4756590420a92cc7))


### Bug Fixes

* **data:** resolve provider merge bug ([b9f38f0](https://github.com/scottmckendry/mnemstart/commit/b9f38f0f2d37457853d21b671ebe78cffd365cb1))
* **keymaps:** sanitize keymap urls before navigating ([a9a389b](https://github.com/scottmckendry/mnemstart/commit/a9a389bac364f53f6c87157399d9e1ff9eb14ab2))

## [0.3.0](https://github.com/scottmckendry/mnemstart/compare/v0.2.1...v0.3.0) (2024-08-23)


### Features

* **auth:** add config opt to send port number in callback urls ([9ed14d9](https://github.com/scottmckendry/mnemstart/commit/9ed14d9f640e3ed2fabad28c8b12027996a23065))

## [0.2.1](https://github.com/scottmckendry/mnemstart/compare/v0.2.0...v0.2.1) (2024-08-23)


### Bug Fixes

* **auth:** rovert sessions package to version supported by goth ([b76f922](https://github.com/scottmckendry/mnemstart/commit/b76f92267a0324e7e564064fec22fa795bb26b94))
* **ci:** include latest tag for published packages ([717f455](https://github.com/scottmckendry/mnemstart/commit/717f455c68ab98d28f2fe63b1196341f8b0fa476))

## [0.2.0](https://github.com/scottmckendry/mnemstart/compare/v0.1.3...v0.2.0) (2024-08-23)


### Features

* **ci:** add dockerfile ([3521d47](https://github.com/scottmckendry/mnemstart/commit/3521d47907dab5cbaca198c4d16ee3a165ccad7d))

## [0.1.3](https://github.com/scottmckendry/mnemstart/compare/v0.1.2...v0.1.3) (2024-08-22)


### Bug Fixes

* **ci:** add release please output tag name for image tag ([12c4a9c](https://github.com/scottmckendry/mnemstart/commit/12c4a9cb8e7ca4522f904c11b8c4ccd14c1e143e))

## [0.1.2](https://github.com/scottmckendry/mnemstart/compare/v0.1.1...v0.1.2) (2024-08-22)


### Bug Fixes

* **ci:** add missing env vars ([b08d800](https://github.com/scottmckendry/mnemstart/commit/b08d800fe1cdf619ced0f342127a5a2ed725eb0f))

## [0.1.1](https://github.com/scottmckendry/mnemstart/compare/v0.1.0...v0.1.1) (2024-08-22)


### Bug Fixes

* **ci:** check for create release rather than tag ref ([8c3e7c2](https://github.com/scottmckendry/mnemstart/commit/8c3e7c281c6ceac91acb9c2ec6fceac37c4ae75f))

## 0.1.0 (2024-08-22)


### Features

* add date & time, tidy up styling ([d55adce](https://github.com/scottmckendry/mnemstart/commit/d55adce33ef56d79ca3d9038e88176b0a273f10d))
* add status line ([9a4611e](https://github.com/scottmckendry/mnemstart/commit/9a4611ecf45e67468010dfee28a8f3440db38d76))
* **auth:** add discord authentication provider ([adbe180](https://github.com/scottmckendry/mnemstart/commit/adbe180d60c61a48d41edac2f5aa530506a52f96))
* **auth:** update stored user during oauth2 callback ([4d5d7f2](https://github.com/scottmckendry/mnemstart/commit/4d5d7f2ce59983004044d24cd153e7969bafab96))
* **ci:** add ci workflow file ([3fbf4c7](https://github.com/scottmckendry/mnemstart/commit/3fbf4c71770af5d3c2b6afd74913866a9faf25c6))
* **ci:** enable dependabot ([56a7410](https://github.com/scottmckendry/mnemstart/commit/56a74106740f881f69b852e047a5b2217c86d744))
* **data:** store authenticated users in db ([8027db0](https://github.com/scottmckendry/mnemstart/commit/8027db03195fc2ed5397e5ba4021a46f57ceb213))
* **keymaps:** add all the CRUD stuff for mappings ([777b9ac](https://github.com/scottmckendry/mnemstart/commit/777b9ac90c0d923678d8ad0fa013a39a2ea594bd))
* **keymaps:** introduce 'leader mode' for more robust mappings ([19faab5](https://github.com/scottmckendry/mnemstart/commit/19faab5d05804d0fa715dc003601708a6a3ee2f0))
* **keymaps:** PoC for key event driven shortcuts with client-side JS ([b365d7b](https://github.com/scottmckendry/mnemstart/commit/b365d7b567a797d1f92a73a62d23434936b42a26))
* minimal icon navigation ([ca46608](https://github.com/scottmckendry/mnemstart/commit/ca4660876a8d35cb4d3097fb278a5d4ec29f1b7f))
* **settings:** add basic user settings schema ([2b871a2](https://github.com/scottmckendry/mnemstart/commit/2b871a2af455b9233c4c646a71058ee7e3ab774c))
* **settings:** add settings edit modal and backend logic ([6e1a77b](https://github.com/scottmckendry/mnemstart/commit/6e1a77bb474b38d851fad5982babb59e710bd293))
* **settings:** render HTML with all current user settings present ([08b2c14](https://github.com/scottmckendry/mnemstart/commit/08b2c141193c98f3978dfe13985465e068d31d4c))
* tidy up modal styling ([54db3d0](https://github.com/scottmckendry/mnemstart/commit/54db3d0a4e5863809e8f1d4aa45195270fdd364a))


### Bug Fixes

* **db:** update user name field conditionally ([4d7761d](https://github.com/scottmckendry/mnemstart/commit/4d7761d70c4359112960078fa9cafb512f642cf9))
* **log:** remove redundant logs ([c7bea13](https://github.com/scottmckendry/mnemstart/commit/c7bea13b4e292f7944d7dca08905a4e913a7cbe5))
