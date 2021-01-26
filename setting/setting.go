package setting

import (
	"log"
	"os/exec"
)

func GetActiveNetWorkDevice() string {
	out, err := exec.Command("/bin/zsh","-c","route -n get default | grep interface | awk '{print $2}'").Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	//fmt.Printf("The output is %s\n", out)

	if out[len(out)-1] == '\n'{
		out = out[:len(out)-1]
	}

	servDev:="networksetup -listallhardwareports | grep -B 1 \""+string(out)+"$\" | awk '/Hardware Port/{ print }'|cut -d \" \" -f3-"

	//fmt.Println(servDev)

	out, err = exec.Command("/bin/zsh","-c",servDev).Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	if out[len(out)-1] == '\n'{
		out = out[:len(out)-1]
	}
	//fmt.Println("The active device is", string(out))
	return string(out)
}

func SetProxy(mode int)  {

	pacproxy:= "http://127.0.0.1:50211/gfw.js"

	activeNetDev:=GetActiveNetWorkDevice()
	if activeNetDev == ""{
		log.Println("can't get active device")
		return
	}

	if mode == 0{
		cmdline:= "networksetup -setsocksfirewallproxystate  "+activeNetDev+" off"
		out, err := exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("turn off sock proxy failed")
			return
		}else{
			log.Println("turn off sock proxy success",string(out))
		}

		cmdline="networksetup -setautoproxyurl \""+ activeNetDev + "\" \""+pacproxy+"\""
		out, err = exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("set auto proxy failed")
			return
		}else{
			log.Println("set auto proxy success",string(out))
		}

		cmdline = "networksetup -setautoproxystate "+activeNetDev+" on"
		out, err = exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("turn on auto proxy failed")
			return
		}else{
			log.Println("turn on auto proxy success",string(out))
		}
	}

	if mode == 1{
		cmdline := "networksetup -setautoproxystate "+activeNetDev+" off"
		out, err := exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("turn off auto proxy failed")
			return
		}else{
			log.Println("turn off auto proxy success",string(out))
		}

		cmdline ="networksetup -setsocksfirewallproxy "+activeNetDev+" 127.0.0.1 50212"
		out, err = exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("set sock proxy failed")
			return
		}else{
			log.Println("set sock proxy success",string(out))
		}

		cmdline = "networksetup -setsocksfirewallproxystate  "+activeNetDev+" on"
		out, err = exec.Command("/bin/zsh","-c",cmdline).Output()
		if err != nil {
			log.Println("turn off sock proxy failed")
			return
		}else{
			log.Println("turn on sock proxy success",string(out))
		}
	}

}

func ClearProxy()  {
	activeNetDev:=GetActiveNetWorkDevice()
	if activeNetDev == ""{
		log.Println("can't get active device")
		return
	}

	cmdline := "networksetup -setautoproxystate "+activeNetDev+" off"
	out, err := exec.Command("/bin/zsh","-c",cmdline).Output()
	if err != nil {
		log.Println("turn off auto proxy failed")
		return
	}else{
		log.Println("turn off auto proxy success",string(out))
	}
	cmdline= "networksetup -setsocksfirewallproxystate  "+activeNetDev+" off"
	out, err = exec.Command("/bin/zsh","-c",cmdline).Output()
	if err != nil {
		log.Println("turn off sock proxy failed")
		return
	}else{
		log.Println("turn off sock proxy success",string(out))
	}

}
