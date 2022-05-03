// 树状导航的处理函数

// 对普通的数组数据进行处理，转换为element tree可以识别的数据
export function DataProcess (data) {
  let trees = []
  if (data != null) {
    // 先遍历第一层
    for (let i = 0; i < data.length; i++) {
      // 第一层的parent为0，我们把这些都放到一个数组里去
      if (data[i].parent === 0) {
        trees.push({ id: data[i].id, label: data[i].title })
      }
    }
    // 然后我们再把子节点插入到父节点里去
    trees = InsertChild(trees, data)
  }
  return trees
}

// 无限遍历插值
function InsertChild (trees, data) {
  for (let i = 0; i < trees.length; i++) {
    for (let j = 0; j < data.length; j++) {
      // 如果这个文档的父节点为父文档id，那么就进行插值操作
      if (trees[i].id === data[j].parent) {
        if (Object.prototype.hasOwnProperty.call(trees[i], 'children')) {
          trees[i].children.push({ id: data[j].id, label: data[j].title })
        } else {
          trees[i].children = [{ id: data[j].id, label: data[j].title }]
        }
      }
    }
    // 同样我们对子节点进行操作
    if (Object.prototype.hasOwnProperty.call(trees[i], 'children')) {
      trees[i].children = InsertChild(trees[i].children, data)
    }
  }
  return trees
}
