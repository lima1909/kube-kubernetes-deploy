package main

import (
	"k8s.io/kubernetes/pkg/kubectl/cmd"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func main() {

	cf := genericclioptions.NewConfigFlags()
	factory := util.NewFactory(cf)
	c := cmd.NewCmdApply("kubectl", factory, genericclioptions.NewTestIOStreamsDiscard())

	c.Flags().Set("filename", "pod.yml")
	err := c.Execute()
	if err != nil {
		panic(err)
	}
}
