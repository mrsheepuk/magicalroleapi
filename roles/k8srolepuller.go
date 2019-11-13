package roles

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func ByName(name string) (error, string[]) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	roles, err := clientset.RbacV1().RolesGetter.List()
	if err != nil {
		panic(err.Error())
	}
	return nil, roles
}

func ByNames(names []string) (error, string) {

}