// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package remove

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/v3/api/filesys"
	"sigs.k8s.io/kustomize/v3/api/ifc"
)

// NewCmdRemove returns an instance of 'remove' subcommand.
func NewCmdRemove(
	fSys filesys.FileSystem,
	v ifc.Validator) *cobra.Command {
	c := &cobra.Command{
		Use:   "remove",
		Short: "Removes items from the kustomization file.",
		Long:  "",
		Example: `
	# Removes resources from the kustomization file
	kustomize edit remove resource {filepath} {filepath}
	kustomize edit remove resource {pattern}

	# Removes one or more patches from the kustomization file
	kustomize edit remove patch <filepath>

	# Removes one or more commonLabels from the kustomization file
	kustomize edit remove label {labelKey1},{labelKey2}

	# Removes one or more commonAnnotations from the kustomization file
	kustomize edit remove annotation {annotationKey1},{annotationKey2}
`,
		Args: cobra.MinimumNArgs(1),
	}
	c.AddCommand(
		newCmdRemoveResource(fSys),
		newCmdRemoveLabel(fSys, v.MakeLabelNameValidator()),
		newCmdRemoveAnnotation(fSys, v.MakeAnnotationNameValidator()),
		newCmdRemovePatch(fSys),
	)
	return c
}
