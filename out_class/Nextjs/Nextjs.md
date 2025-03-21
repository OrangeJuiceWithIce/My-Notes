# Next.js

## TailWind CSS
相当于把css嵌入到className中
```js
<div className='text-red-500'>This is a red text</div>
```
## 优化font和image
先获取字体
```js
import {Inter} from 'next/font/google';

export const inter=Inter({subsets:['latin']});
```
再在className中使用字体的className
```js
<body className={'${inter.className} antialiased'}>{children}</body>
```
## 预渲染

