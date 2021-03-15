package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts INTEGER_ARRAY nodeValues as parameter.
 */
 
 type Node struct {
     value int
     left *Node
     right *Node
    //  parent *Node
    //  kind string // l: left, r: right, t: top
 }

func isValid(nodeValues []int32) bool {
    // number of processed nodes
    //count := 0
    if len(nodeValues) == 0 {
        return false
    }


    maxDepth := 7
    activeNodes := make([]int32, maxDepth)
    activeNodes[0] = nodeValues[0]
    depth := 1

    valid, ofs := processLeft(nodeValues, &activeNodes, 1, depth)
    if !valid {
        return false;
    }
    valid, _ = processRight(nodeValues, &activeNodes, ofs, depth)
    return valid
}

func processLeft(nodeValues []int32, activeNodes *[]int32, ofs int, depth int) (bool, int) {
    if len(nodeValues) == ofs {
        return true, ofs
    }
    if nodeValues[ofs] >= nodeValues[ofs-1] {
        return true, ofs
    }

    (*activeNodes)[depth] = nodeValues[ofs]
    ofs++
    depth++

    valid, ofs := processLeft(nodeValues, activeNodes, ofs, depth)
    if valid {
        valid, ofs = processRight(nodeValues, activeNodes, ofs, depth)
    }
    return valid, ofs
}

func processRight(nodeValues []int32, activeNodes *[]int32, ofs int, depth int) (bool, int) {
    if len(nodeValues) == ofs {
        return true, ofs
    }
    if nodeValues[ofs] < (*activeNodes)[depth-1] {
        return false, ofs
    }

    (*activeNodes)[depth] = nodeValues[ofs]
    ofs++
    depth++

    // None of the nodes that are currently active can be greater than the value being processed.
    if len(nodeValues) == ofs {
        return true, ofs
    }
    for i := 0; i < depth-1; i++ {
        if nodeValues[ofs] < (*activeNodes)[i] {
            return false, ofs
        }
    }

    valid, ofs := processLeft(nodeValues, activeNodes, ofs, depth)
    if valid {
        valid, ofs = processRight(nodeValues, activeNodes, ofs, depth)
    }
    return valid, ofs
    
}




func isValid1(nodeValues []int32) bool {
   all := make([]int, 0, len(nodeValues))
    for _, el := range nodeValues {
        all = append(all, int(el))
    }
    
    //todo: compare not sorted
    sort.Ints(all)
    
    root := arrayToTree(all, 0, len(nodeValues) - 1, int(nodeValues[0]))
    
    // Write your code here
    // root := &Node{
    //     value: nodeValues[0],
    //     kind: "t",
    //     }
    // current := root
    // count++
    // for count < len(nodeValues) {
    //     var ok bool
    //     var left *Node
    //     var right *Node
    //     if (count + 1) < len(nodeValues) {
    //         left, right, ok = fitChildren(current, nodeValues[count], nodeValues[count+1])
    //     } else {
    //         left, right, ok = fitChildren(current, nodeValues[count], -1)
    //     }
    //     fmt.Printf("left: %v; right: %v; ok:%v\n", left, right, ok)
    //     if !ok {
    //         return false
    //     }
    //     if left != nil {
    //         count++
    //     }
    //     if right != nil {
    //         count++
    //     }
    // }
    inOrder := traverseInOrder(root)
    last := inOrder[0]
    fmt.Printf("\n in order printing: \n")
    for _, el := range inOrder {
        if last > el {
            return false;
        }
        fmt.Printf("node: %v\n", el)
    }
    return true
}

func arrayToTree(array []int, fromIdx int, toIdx int, centerVal int) *Node {
    if fromIdx > toIdx {
        return nil
    }
    var middle int
    if centerVal != -1 {
        middle = getIndex(array, centerVal)
    } else {
        middle = (fromIdx + toIdx) / 2
    }
    node := &Node{
        value: array[middle],
    }
    node.left = arrayToTree(array, fromIdx, middle - 1, -1)
    node.right = arrayToTree(array, middle + 1, toIdx, -1)
    return node
}

func getIndex(array []int, val int) int {
    for idx, el := range array {
        if el == val {
            return idx
        }
    }
    return -1
}

// // returns: left, right, ok
// func fitChildren(node *Node, leftCandidate int32, rightCandidate int32) ( *Node,  *Node,  bool) {
//     fmt.Printf("fit: %v %v %v\n", node.value, leftCandidate, rightCandidate)
//     if leftCandidate == node.value {
//         // not expected
//         fmt.Printf("error: 2 nodes with the same value")
//         //panic("error: 2 nodes with the same value")
//         return nil, nil, false
//     }
//     var left *Node
//     var right *Node
//     if leftCandidate < node.value {
//         // fits for the left
//         left = &Node{
//             value: leftCandidate,
//             parent: node,
//             kind: "l",
//             }
//         node.left = left
//     } else {
//         // fits for the right
//         right = &Node{
//             value: leftCandidate,
//             // parent: node,
//             // kind: "r",
//             }
//         node.right = right
//         return nil, right, true
//     }
//     // last item
//     if rightCandidate == -1 {
//         return left, nil, true
//     }
    
//     if rightCandidate > node.value {
//         // fits for the right
//         right = &Node{value: rightCandidate}
//         // right.parent = node
//         // right.kind = "r"
//         node.right = right
//     } else {
//         return left, nil, true
//     }
//     return left, right, true     
// }

func traverseInOrder(root *Node) []int {
    stack := []int{}
    return subTree(root, stack)
}

func subTree(node *Node, stack []int) []int {
    if node == nil {
        return stack
    }
    if node.left != nil {
        stack = subTree(node.left, stack)
    }
    stack = append(stack, node.value)
    if node.right != nil {
        stack = subTree(node.right, stack)
    }
    return stack
}
 
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)

        aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var a []int32

        for i := 0; i < int(aCount); i++ {
            aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
            checkError(err)
            aItem := int32(aItemTemp)
            a = append(a, aItem)
        }

        var result string
        if (isValid(a)) {
            result = "1"
        } else {
            result = "0"
        }

        fmt.Fprintf(writer, "%s\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}