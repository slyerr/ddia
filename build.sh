#!/bin/bash

# 清理
rm -rf build/temp
rm -f build/数据密集型应用系统设计*
# 生成 .epub
gitbook epub ./ build/数据密集型应用系统设计.epub
ebook-meta build/数据密集型应用系统设计.epub --author-sort 'Kleppmann, Martin' --date '' --publisher '' --rating '5' --book-producer ''
unzip -o build/数据密集型应用系统设计.epub -d build/temp
sed -i 's/目錄/目录/g' build/temp/SUMMARY.html
sed -i 's/serif/Source Han Serif CN/g' build/temp/SUMMARY.html
cd build/temp
zip -q -r 数据密集型应用系统设计.epub *
cd ../../

file=数据密集型应用系统设计_$(date +%Y%m%d%H%M)
# 设置字体
cp build/temp/数据密集型应用系统设计.epub build/$file.epub
ebook-convert build/temp/数据密集型应用系统设计.epub build/$file.mobi 

rm -rf build/temp
rm -f build/数据密集型应用系统设计.*