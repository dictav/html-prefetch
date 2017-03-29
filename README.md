HTML Prefetch
=============

- [Link prefetching FAQ](https://developer.mozilla.org/en-US/docs/Web/HTTP/Link_prefetching_FAQ)
- [Source Hints](https://www.w3.org/TR/resource-hints/)

2017/03/29 現在、Chrome などのモダンなブラウザではページ遷移を高速化するために Preload 機構を導入している。これは、<link> タグや Link ヘッダーにヒントを書くことでブラウザがIDLE時間を使って遷移先のページやアセットを先読みすることでページ遷移を高速化します。

## 対応状況

* TN: `<link rel="next" href="/prefetch.html">` タグがある場合
* TPF: `<link rel="prefetch" href="/prefetch.html">` タグがある場合
* TPR: `<link rel="prerender" href="/prefetch.html">` タグがある場合
* HN: `Link: </prefetch.html>; rel=next` ヘッダーがある場合
* HPF: `Link: </prefetch.html>; rel=prefetch` ヘッダーがある場合
* HPR: `Link: </prefetch.html>; rel=prerender` ヘッダーがある場合
* PRE: リンクを長押ししてプレビューを表示する場合

| OS            | Browser     | HN | HPF | HPR | TN | TPF | TPR | PRE | note                              |
|---------------|-------------|----|-----|-----|----|-----|-----|-----|-----------------------------------|
| macOS 10.12.4 | Firefox 52  | ✔︎  | ✔︎   | ✖︎   | ✔︎  | ✔︎   | ✖︎   | -   | Prefetch時にX-mozを送信する       |
| macOS 10.12.4 | Chrome 57   | ✖︎  | ✔︎   | ✖︎   | ✖︎  | ✔︎   | ✔︎   | -   | TPR時にPrefetchしたJSが実行される |
| macOS 10.12.4 | Safari 10.2 | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✖︎   | ✖︎   | ✔︎   | PRE時にPrefetchしたJSが実行される |
| iOS 10        | Safari      | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✖︎   | ✖︎   | ?   |                                   |
| Android 5.0   | Firefox 52  | ✔︎  | ✔︎   | ✖︎   | ✔︎  | ✔︎   | ✖︎   | -   | Prefetch時にX-mozを送信する       |
| Android 5.0   | Chrome 56   | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✔︎   | ✔︎   | -   | TPR時にPrefetchしたJSが実行される |
| Windows Vista | IE 7        | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✖︎   | ✖︎   | -   |                                   |
| Windows 7     | IE 8        | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✖︎   | ✖︎   | -   |                                   |
| Windows 8.1   | IE 11       | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✔︎   | ✔︎   | -   | TPR時にPrefetchしたJSが実行される |
| Windows 10    | Edge 14     | ✖︎  | ✖︎   | ✖︎   | ✖︎  | ✖︎   | ✖︎   | -   |                                   |

## Build beacon.js

## Test

```
# console1
$ go run main.go
```

```
# console2
$ open http://localhost:8080
```

各ページにアクセスすると `console1` にログが出力されます。
また、各ページは読み込み時に Beacon を飛ばしており、実行されると `echo: Header Next` のようなログが出力されます。
さらに、Beacon を送信時に `document.visibilityState !== 'visible'` が検出された場合は `log: prerender` のようなログが出力されます。

## 注意

- ブラウザがファイルをキャッシュしているときは Prefetch は行われない
- PrefetchしたJSを実行するのは Chrome と Safari。ただしSafariは実際に表示されているので Chrome と等価ではない。
- Chrome で EventListener `visibilitychange` を追加することで、`visibilityState === 'visible'` になった時に Beacon を飛ばすことが可能
