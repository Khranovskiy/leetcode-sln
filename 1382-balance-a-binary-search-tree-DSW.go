package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Step struct {
	description string
	tree        *TreeNode
}

// Bubble Tea Model
type model struct {
	steps       []Step
	currentStep int
	width       int
	height      int
	quitting    bool
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)

	stepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	treeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA"))

	navStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)

	highlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF00FF")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500"))
)

func initialModel(tree *TreeNode) model {
	steps := generateSteps(tree)
	return model{
		steps:       steps,
		currentStep: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "right", "l", "n", " ", "enter":
			if m.currentStep < len(m.steps)-1 {
				m.currentStep++
			}

		case "left", "h", "p":
			if m.currentStep > 0 {
				m.currentStep--
			}

		case "home":
			m.currentStep = 0

		case "end":
			m.currentStep = len(m.steps) - 1

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num, _ := strconv.Atoi(msg.String())
			if num >= 1 && num <= len(m.steps) {
				m.currentStep = num - 1
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return highlightStyle.Render("\n  Thanks for exploring DSW algorithm! 👋\n\n")
	}

	step := m.steps[m.currentStep]

	var b strings.Builder

	// Title
	title := fmt.Sprintf(" DSW Algorithm Visualizer - Step %d/%d ", m.currentStep+1, len(m.steps))
	b.WriteString(titleStyle.Render(title))
	b.WriteString("\n\n")

	// Step description
	b.WriteString(stepStyle.Render(step.description))
	b.WriteString("\n\n")

	// Tree visualization
	if step.tree != nil {
		treeStr := renderTree(step.tree)
		b.WriteString(treeStyle.Render(treeStr))
		b.WriteString("\n")

		// In-order traversal
		inorder := infoStyle.Render("In-order: ") + getInOrder(step.tree)
		b.WriteString(inorder)
		b.WriteString("\n\n")
	}

	// Progress bar
	progress := renderProgressBar(m.currentStep, len(m.steps), 50)
	b.WriteString(progress)
	b.WriteString("\n\n")

	// Navigation help
	nav := navStyle.Render("Navigation: ")
	if m.currentStep > 0 {
		nav += highlightStyle.Render("← ") + navStyle.Render("prev  ")
	}
	if m.currentStep < len(m.steps)-1 {
		nav += highlightStyle.Render("→ ") + navStyle.Render("next  ")
	}
	nav += highlightStyle.Render("home ") + navStyle.Render("first  ")
	nav += highlightStyle.Render("end ") + navStyle.Render("last  ")
	nav += highlightStyle.Render("q ") + navStyle.Render("quit")
	b.WriteString(nav)

	return b.String()
}

func renderProgressBar(current, total, width int) string {
	if total == 0 {
		return ""
	}

	percentage := float64(current) / float64(total-1)
	filled := int(percentage * float64(width))

	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else if i == filled {
			bar += "▓"
		} else {
			bar += "░"
		}
	}
	bar += fmt.Sprintf("] %d%%", int(percentage*100))

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Render(bar)
}

func renderTree(root *TreeNode) string {
	var b strings.Builder
	printTreeToBuilder(&b, root, "", false)
	return b.String()
}

func printTreeToBuilder(b *strings.Builder, root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	connector := "└── "
	if isLeft {
		connector = "├── "
	}
	b.WriteString(prefix + connector + fmt.Sprint(root.Val) + "\n")

	if root.Left != nil || root.Right != nil {
		extension := "    "
		if isLeft {
			extension = "│   "
		}
		if root.Left != nil {
			printTreeToBuilder(b, root.Left, prefix+extension, true)
		}
		if root.Right != nil {
			printTreeToBuilder(b, root.Right, prefix+extension, false)
		}
	}
}

func getInOrder(root *TreeNode) string {
	var values []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	var b strings.Builder
	for i, v := range values {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprint(v))
	}
	return b.String()
}

// Tree building and parsing functions
func buildTreeFromString(s string) *TreeNode {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	if s == "" {
		return nil
	}

	parts := strings.Split(s, ",")
	values := make([]*int, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "null" {
			values = append(values, nil)
		} else {
			val, err := strconv.Atoi(part)
			if err != nil {
				panic(fmt.Sprintf("Invalid number: %s", part))
			}
			values = append(values, &val)
		}
	}

	return buildTree(values)
}

func buildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  cloneTree(root.Left),
		Right: cloneTree(root.Right),
	}
}

// DSW Algorithm with step recording
func generateSteps(root *TreeNode) []Step {
	steps := []Step{}

	if root == nil {
		return steps
	}

	steps = append(steps, Step{
		description: "Step 0: Original unbalanced BST",
		tree:        cloneTree(root),
	})

	dummy := &TreeNode{Right: root}
	vine := createVine(dummy, &steps)
	compress(dummy, vine, &steps)

	steps = append(steps, Step{
		description: fmt.Sprintf("✓ Final: Balanced BST created! Total steps: %d", len(steps)),
		tree:        cloneTree(dummy.Right),
	})

	return steps
}

func createVine(grand *TreeNode, steps *[]Step) int {
	count := 0
	node := grand.Right
	stepNum := 1

	for node != nil {
		if node.Left != nil {
			oldNode := node
			node = node.Left
			grand.Right = node
			oldNode.Left = node.Right
			node.Right = oldNode

			*steps = append(*steps, Step{
				description: fmt.Sprintf("Right rotation at node %d (creating vine)", oldNode.Val),
				tree:        cloneTree(grand.Right),
			})
			stepNum++
		} else {
			count++
			grand = node
			node = node.Right
		}
	}

	*steps = append(*steps, Step{
		description: fmt.Sprintf("✓ Vine created: %d nodes in right-skewed list", count),
		tree:        cloneTree((*steps)[len(*steps)-1].tree),
	})

	return count
}

func compress(grand *TreeNode, count int, steps *[]Step) {
	h := 0
	temp := count + 1
	for temp > 1 {
		temp >>= 1
		h++
	}
	m := (1 << h) - 1

	*steps = append(*steps, Step{
		description: fmt.Sprintf("Starting compression: %d nodes → target perfect tree size: %d", count, m),
		tree:        cloneTree(grand.Right),
	})

	makeCompression(grand, count-m, steps, "bottom level")

	compNum := 1
	for m > 1 {
		m /= 2
		makeCompression(grand, m, steps, fmt.Sprintf("level %d", compNum))
		compNum++
	}
}
func makeCompression(grand *TreeNode, m int, steps *[]Step, level string) {
	if m <= 0 {
		return
	}

	node := grand.Right
	for i := 0; i < m; i++ {
		if node == nil {
			break
		}
		oldNode := node
		node = node.Right
		grand.Right = node
		if node != nil {
			oldNode.Right = node.Left
			node.Left = oldNode
			grand = node
			node = node.Right

			*steps = append(*steps, Step{
				description: fmt.Sprintf("Compression %s: left rotation %d/%d at node %d", level, i+1, m, oldNode.Val),
				tree:        cloneTree(grand.Right),
			})
		}
	}
}

func selectTestCase() string {
	testCases := map[string]string{
		"1": "[1,null,2,null,3,null,4,null,5,null,6,null,7]",
		"2": "[1,null,2,null,3,null,4]",
		"3": "[10,5,15,2,7,null,20,1,null,6,8]",
		"4": "[4,2,6,1,3,5,7]",
	}

	fmt.Println(titleStyle.Render(" Day-Stout-Warren Algorithm Visualizer "))
	fmt.Println()
	fmt.Println("Choose a test case:")
	fmt.Println("  1) Extreme right skew: [1,null,2,null,3,null,4,null,5,null,6,null,7]")
	fmt.Println("  2) Simple right skew:  [1,null,2,null,3,null,4]")
	fmt.Println("  3) Mixed tree:         [10,5,15,2,7,null,20,1,null,6,8]")
	fmt.Println("  4) Already balanced:   [4,2,6,1,3,5,7]")
	fmt.Println()
	fmt.Print("Enter choice (1-4): ")

	var choice string
	fmt.Scanln(&choice)

	treeStr, ok := testCases[choice]
	if !ok {
		treeStr = testCases["1"]
	}

	return treeStr
}

func main() {
	treeStr := selectTestCase()
	tree := buildTreeFromString(treeStr)

	p := tea.NewProgram(initialModel(tree), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
