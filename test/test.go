package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	test1()
}

func test1()  {
	//cmdline:="/bin/zsh \"\""

	out, err := exec.Command("/bin/zsh","-c","route get default | grep interface | awk '{print $2}'").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The output is %s\n", out)

	if out[len(out)-1] == '\n'{
		out = out[:len(out)-1]
	}

	servDev:="networksetup -listallhardwareports | grep -B 1 "+string(out)+" | awk '/Hardware Port/{ print }'|cut -d \" \" -f3-"

	fmt.Println(servDev)

	out, err = exec.Command("/bin/zsh","-c",servDev).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	if out[len(out)-1] == '\n'{
		out = out[:len(out)-1]
	}
	fmt.Println("The active device is", string(out))

}