package main

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
			return true
		} else if p == nil || q == nil || q.Val != p.Val {
			return false
		} else {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right);
		}
}