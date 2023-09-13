var md = new MobileDetect(navigator.userAgent);

const makeNode = (nodeName) => {
  let node = document.createElement(nodeName);
  return node;
};

const appendNode = (nodeName, childNode) => {
  let node = makeNode(nodeName);
  node.appendChild(childNode);
  return node;
};

const textNode = (text) => {
  return document.createTextNode(text);
};

const listNode = (children) => {
  return appendNode("li", appendNode("p", children));
};

const makeList = (text) => {
  let ul = document.getElementById("contents");
  ul.appendChild(listNode(textNode(text)));
};

const viewPort = (content) => {
  var meta = document.createElement("meta");
  meta.name = "viewport";
  meta.content = content;
  document.head.appendChild(meta);
};

window.addEventListener("load", function () {
  if (md.mobile() === null) {
    for (let i = 0; i < 8; i++) {
      makeList("you in PC");
    }
  } else {
    for (let i = 0; i < 10; i++) {
      viewPort("width=device-width,initial-scale=1,user-scalable=no");
      makeList("you in Mobile");
    }
  }
});
