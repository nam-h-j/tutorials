/* eslint-disable @typescript-eslint/no-var-requires */
/* cache clear when do the build */
const replace = require('replace-in-file')
const colors = require('colors')
const crypto = require('crypto')
const fs = require('fs')
const path = require('path')
const hash = crypto.randomBytes(8).toString('hex')

const root = path.resolve('')
const pagesJsPath = path.join(root, '/dist')

const allPageJsFiles = []
function printAllFiles(dir, relativePath) {
  const filenames = fs.readdirSync(dir)
  filenames.forEach((filename) => {
    const file = path.join(relativePath, filename)
    const fullPath = path.join(dir, filename)
    const stats = fs.statSync(fullPath)
    if (stats.isFile() && file.match('.html')) {
      allPageJsFiles.push(`dist${relativePath}${filename}`)
    } else if (stats.isDirectory()) {
      printAllFiles(fullPath, `${relativePath}${filename}/`)
    }
  })
}
printAllFiles(pagesJsPath, '/')

function replaceHashCode(option) {
  replace(option)
    .then((files) => {
      const changedFiles = []
      files.forEach((file) => {
        if (file.hasChanged) {
          changedFiles.push(file.file)
        }
      })
      console.log(colors.green('Cache clear in pages js files. {HASH} \n'), changedFiles)
    })
    .catch((error) => {
      console.error('Error occurred:', error)
    })
}

const allFilesOptions = {
  files: allPageJsFiles,
  from: /.css"/g,
  to: `.css?h=${hash}"`,
}
replaceHashCode(allFilesOptions)
