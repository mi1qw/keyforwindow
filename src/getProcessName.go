package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <proc/readproc.h>

char* getProcessName(int pid) {
    char* command_line = NULL;

    PROCTAB* proc = openproc(PROC_FILLSTATUS);
    if (proc == NULL) {
        return NULL;
    }

    proc_t proc_info;
    memset(&proc_info, 0, sizeof(proc_info));

    while (readproc(proc, &proc_info) != NULL) {
        if (proc_info.tid == pid) {
            command_line = (char*)malloc(strlen(proc_info.cmdline)+1);
            strcpy(command_line, proc_info.cmdline);
            break;
        }
    }

    closeproc(proc);
    return command_line;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func GetProcessName(pid int) (string, error) {
	result := C.getProcessName(C.int(pid))
	if result == nil {
		return "", fmt.Errorf("cannot get process name for pid %v", pid)
	}
	defer C.free(unsafe.Pointer(result))

	return C.GoString(result), nil
}

func main() {
	pid := 15408 // ваш ID процесса
	processName, err := GetProcessName(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Имя процесса с PID %d: %s\n", pid, processName)
}
