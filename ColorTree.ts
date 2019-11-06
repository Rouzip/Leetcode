enum Color {
  black = 0,
  white = 1
}

class MyNode {
  color: Color;
  child: MyNode[];

  constructor(childList: MyNode[]) {
    this.color = 1;
    this.child = childList;
  }
}

function blackNodeNumber(node: MyNode): number {
  let result: number = 0;
  const changeColor: Function = function(node: MyNode): void {
    if (node === null) {
      return;
    }
    if (node.child.length > 0) {
      for (let i = 0; i < node.child.length; i++) {
        let childNode = node.child[i];
        childNode.color = 0;
        result++;
        changeColor(childNode);
      }
    }
    node.color = 1;
    result--;
  };
  const changeColor2: Function = function(node: MyNode): void {
    if (node === null || node.child.length == 0 || node.color == 0) {
      return;
    }

    for (let i = 0; i < node.child.length; i++) {
      let childNode = node.child[i];
      let tmp = false;

      for (let j = 0; j < childNode.child.length; j++) {
        if (childNode.child[j].color == 0) {
          tmp = true;
        }
      }

      if (!tmp) {
        childNode.color = 0;
        result++;
      }
    }

    node.child.map(childNode => {
      changeColor2(childNode);
    });
  };

  if (node == null) {
    return 0;
  }

  changeColor(node);
  changeColor2(node);

  return result;
}

let root = new MyNode([]);
let l1 = new MyNode([
  new MyNode([]),
  new MyNode([]),
  new MyNode([]),
  new MyNode([]),
  new MyNode([])
]);
let l2 = new MyNode([new MyNode([]), new MyNode([]), new MyNode([])]);

let l3 = new MyNode([new MyNode([]), new MyNode([])]);

root.child.push(l1, l2, l3);

console.log(blackNodeNumber(root));
