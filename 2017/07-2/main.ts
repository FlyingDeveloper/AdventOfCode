import fs = require('fs');

class MyNode {
    name: string;
    children: Array<MyNode>;
    parent: MyNode | undefined;
    weight: number;

    constructor() {
        this.children = [];
    }

    public getWeight(): number {
        let rollingSum = this.weight;
        if (this.children && this.children.length > 0) {
            rollingSum += this.children.reduce((aggregator, current) => {
                return aggregator + current.getWeight();
            }, rollingSum);
        }

        return rollingSum;
    }

    public isBalanced(): boolean {
        let expectedWeight: number|undefined = undefined;
        for (var i = 0; i < this.children.length; i++) {
            if (expectedWeight === undefined) {
                expectedWeight = this.children[i].getWeight();
            } else {
                return this.children[i].getWeight() === expectedWeight;
            }
        }

        return true;
    }
}

function findRoot(node: MyNode): MyNode {
    let currentNode:MyNode|undefined = node;

    while (currentNode.parent !== undefined) {
        currentNode = currentNode.parent;
    }

    return currentNode;
}

function findChildInTree(root: MyNode, childToSearchFor: string): MyNode | undefined {
    if (root.name === childToSearchFor) {
        return root;
    }

    for (var i = 0; i < root.children.length; i++) {
        const child = root.children[i];
        if (child.name === childToSearchFor) {
            return child;
        } else {
            if (child.children && child.children.length > 0) {
                let result = findChildInTree(child, childToSearchFor);
                if (result) {
                    return result;
                }

                return undefined;
            } else {
                return undefined;
            }
        }
    };
    
    return undefined;
}

function processInput(input: string) {
    const childlessLineRegex = /^(\w+)\s\((\d+)\)$/;
    const parentLineRegex = /^(\w+)\s\((\d+)\)\s->\s(.*)$/;

    const lines = input.split('\n');
    let root: MyNode;
    let lostChildren = new Map<string, MyNode>(); // map of unknown child names to parent nodes waiting for those children
    let orphans = new Map<string, MyNode>();
    lines.forEach(line => {
        console.log('-----');
        const childlessResult = childlessLineRegex.exec(line);
        const parentResult = parentLineRegex.exec(line);
        let currentNode: MyNode;
        let parent: MyNode | undefined = undefined;

        if (line.length === 0) {
            return;
        }

        if (childlessResult !== null) {
            currentNode = new MyNode();
            currentNode.name = childlessResult[1];
            currentNode.parent = undefined;
            currentNode.weight = Number(childlessResult[2]);

            if (parent = lostChildren.get(currentNode.name)) {
                // Need to add this node to the parent's children
                parent.children.push(currentNode);
                lostChildren.delete(currentNode.name);
                currentNode.parent = parent;
            } else {
                orphans.set(currentNode.name, currentNode);
            }
        } else if (parentResult !== null) {
            currentNode = new MyNode();
            currentNode.name = parentResult[1];
            currentNode.parent = undefined;
            currentNode.weight = Number(parentResult[2]);

            let parentFoundOrSetToRoot = false;
            
            if (parent = lostChildren.get(currentNode.name)) {
                parent.children.push(currentNode);
                lostChildren.delete(currentNode.name);
                parentFoundOrSetToRoot = true;
                currentNode.parent = parent;
            } else if (root === undefined) {
                root = findRoot(currentNode);
                parentFoundOrSetToRoot = true;
            }

            const childNames = parentResult[3].split(',').map(x => x.trim());
            childNames.forEach(child => {
                // Check orphans
                const orphan = orphans.get(child);
                if (orphan) {
                    currentNode.children.push(orphan);
                    orphan.parent = currentNode;
                    if (orphan === root) {
                        root = findRoot(currentNode);
                        parentFoundOrSetToRoot = true;
                    }
                    orphans.delete(child);
                } else {
                    // Check tree starting at root
                    const foundChild: MyNode | undefined = findChildInTree(root, child);
                    if (foundChild !== undefined) {
                        currentNode.children.push(foundChild);
                        if (foundChild === root) {
                            root = findRoot(currentNode);
                            parentFoundOrSetToRoot = true;
                        }
                    } else {
                        // Add unfound children to lostChildren
                        lostChildren.set(child, currentNode);
                    }
                }
            });

            if (!parentFoundOrSetToRoot) {
                orphans.set(currentNode.name, currentNode);
            }
        }

        // Make sure that none of the orphans have the root as a child...if they do, update root
        /*let allOrphans = orphans.values();
        let current: MyNode;
        while (current = allOrphans.next().value) {
            current.children.forEach(currentChild => {
                if (currentChild === root) {
                    root = findRoot(currentNode);
                }
            });
        }*/

        /*console.log(`Orphan count: ${orphans.size}`);
        console.log(`Lost children count: ${lostChildren.size}`);*/
        console.log(`Root: ${root && root.name} - ${root && root.getWeight()} - ${root && root.isBalanced()}`);
    });
}

/* Test Case */
/*processInput(`
pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`); // tknk*/

fs.readFile('input.txt', { encoding: 'utf8' }, (err, contents) => processInput(contents));
//ddspu was wrong
//fhxpkd was wrong
//hmvwl is correct...need to figure out where it has wound up