package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config") // Get kubeconfig

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	envConfig, err := os.Open("envs.txt") // open envConfig file
	if err != nil {
		log.Fatalf("Fatal error: %s \n", err)
	}
	defer envConfig.Close()

	scanner := bufio.NewScanner(envConfig) // create a new scanner over the file

	for scanner.Scan() { // Iterate each line to parse env, cluster and namespace
		lineArray := strings.Split(scanner.Text(), ":")
		env := lineArray[0]
		cluster := lineArray[1]
		ns := lineArray[2]

		fmt.Printf("ENV: %s   CLUSTER:%s   NS:%s\n", env, cluster, ns)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	/*
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for _, b := range pods.Items {
			fmt.Println(b)
	*/

}
