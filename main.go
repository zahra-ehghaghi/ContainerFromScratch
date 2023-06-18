package main
import (
    "fmt"
	"os"
    "os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default: 
	  panic("help")
		
	}
}
func run()  {
	fmt.Printf("Runing %v\n ",os.Args[2:])	
	cmd := exec.Command("/proc/self/exe",append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS| syscall.CLONE_NEWPID|syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	
	err := cmd.Run()
	if err != nil {
		panic("error runing command")
	}
}


func child(){
	fmt.Printf("Runing %v as %d\n",os.Args[2:],os.Getegid())
	syscall.Sethostname([]byte("Devopshobbies "))
	syscall.Chroot("/tmp/alpine-minirootfs/")
	syscall.Chdir("/")
	syscall.Mount("proc","proc","proc",0,"")
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, 
	}
	err := cmd.Run()
	if err != nil {
		panic("error runing command")
	}

}

