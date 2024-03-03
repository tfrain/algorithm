func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    // Declare a slice to hold subtrees that are duplicates
    var subtrees []*TreeNode
    // Initialize a map to keep track of the count of each serialized subtree (in string form)
    cnt := make(map[string]int, 0)
    // find is a helper function to serialize a tree rooted at 'root'
    var find func(root *TreeNode) []byte
    // Define the find function
    find = func(root *TreeNode) []byte {
        // The base case: when the current node is nil, encode it as '#'
        if root == nil {
            return []byte{'#'}
        }
        // Build the serialization string for the subtree rooted at current node
        // Using pre-order traversal: root -> left subtree -> right subtree
        str := append([]byte{byte(root.Val)}, ',')
        // Recursively process the left subtree, and append the returned results to str
        str = append(str, find(root.Left)...) 
        // Append a separator
        str = append(str, ',')
        // Recursively process the right subtree, and append the returned results to str
        str = append(str, find(root.Right)...) 
        // Update the count of current subtree in the map
        cnt[string(str)]++
        // If this subtree has appeared once before (meaning it is a duplicate), add its root to the result list
        if cnt[string(str)] == 2 {
            subtrees = append(subtrees, root)
        }
        // Return the serialization string of current subtree
        return str
    }
    // Start the process from the root of given tree
    find(root)
    // Return all roots of the duplicate subtrees
    return subtrees
}