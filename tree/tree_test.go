package tree

import (
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("makes a new TreeNode", func(t *testing.T) {
		want := 1
		n := NewTreeNode(want)
		got := n.val
		assertEqual(t, want, got)
		assertEqual(t, len(n.children), 0)
	})
}

func TestAddChild(t *testing.T) {
	t.Run("adds a child node", func(t *testing.T) {
		want := 2
		n := NewTreeNode(1)
		n.AddChild(want)
		got := n.children[0].val
		assertEqual(t, want, got)
		assertEqual(t, len(n.children), 1)
	})
}

func TestContains(t *testing.T) {
	t.Run("returns true if present", func(t *testing.T) {
		want := true
		n := NewTreeNode(1)
		n.AddChild(2)
		got := n.Contains(2)
		if got != want {
			t.Errorf("wanted %t, got %t", want, got)
		}
		assertEqual(t, len(n.children), 1)
	})
	t.Run("returns false if not present", func(t *testing.T) {
		want := false
		n := NewTreeNode(1)
		n.AddChild(2)
		got := n.Contains(3)
		if got != want {
			t.Errorf("wanted %t, got %t", want, got)
		}
		assertEqual(t, len(n.children), 1)
	})
}

func TestGetValue(t *testing.T) {
	t.Run("returns the value on the node", func(t *testing.T) {
		want := 1
		n := NewTreeNode(want)
		got := n.GetValue()
		assertEqual(t, want, got)
	})
}

func TestSetValue(t *testing.T) {
	t.Run("overrides previous value", func(t *testing.T) {
		first := 0
		n := NewTreeNode(first)
		assertEqual(t, first, n.GetValue())
		second := 1
		n.SetValue(second)
		assertEqual(t, second, n.GetValue())
	})
}

func TestGetNode(t *testing.T) {
	t.Run("returns the correct node", func(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		needle := 3
		n.children[0].AddChild(needle)
		err, found := n.GetNode(needle)
    if err != nil {
      t.Errorf("got an unexpected error")
    }
		assertEqual(t, needle, found.GetValue())
	})
}

func TestRemoveNode(t *testing.T) {
  t.Run("removes a node and elevates its children", func(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		n.children[0].AddChild(3)
    n.RemoveNode(1) // 3 should move up to take its place
    assertEqual(t, 3, n.children[0].val)
  })
  t.Run("returns an error if the node does not exist", func(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		n.children[0].AddChild(3)
    err := n.RemoveNode(4) // does not exist
    if err == nil {
      t.Error("expected error, got nil")
    }
    assertEqual(t, 1, n.children[0].val)
    assertEqual(t, 2, len(n.children))
  })
}

func TestRemoveBranchByNode(t *testing.T) {
  t.Run("removes a branch starting at the given node value", func(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		n.children[0].AddChild(3)
    n.RemoveBranchByNode(1) // only 2 should still be a child, 3 is discarded
    assertEqual(t, 2, n.children[0].val)
    assertEqual(t, 1, len(n.children))
  })
  t.Run("returns an error if the node does not exist", func(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		n.children[0].AddChild(3)
    err := n.RemoveBranchByNode(4) // does not exist
    if err == nil {
      t.Error("expected error, got nil")
    }
    assertEqual(t, 1, n.children[0].val)
    assertEqual(t, 2, len(n.children))
  })
}

func TestPrint(t *testing.T) {
		n := NewTreeNode(0)
		n.AddChild(1)
		n.AddChild(2)
		n.children[0].AddChild(3)
    n.Print()
}

func assertEqual(t testing.TB, want, got int) {
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
