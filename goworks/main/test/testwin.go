package main
import (
	"fmt"
	"github.com/cwchiu/go-winapi"
	"syscall"
	"unsafe"
)

const(
	HTCLIENT uintptr = 1
	HTCAPTION uintptr = 2
)

func main() {
	ins := winapi.GetModuleHandle(nil)
	pc, _ := syscall.UTF16PtrFromString("文字")
	pt, _ := syscall.UTF16PtrFromString("タイトル")
	
	wc := winapi.WNDCLASSEX{
		CbSize        : uint32(unsafe.Sizeof(winapi.WNDCLASSEX{})),//必須
		Style         : 0,
		LpfnWndProc   : syscall.NewCallback(windowProc),//必須?
		CbClsExtra    : 0,
		CbWndExtra    : 0,
		HInstance     : ins,
		HIcon         : winapi.LoadIcon(ins, winapi.MAKEINTRESOURCE(132)),
		HCursor       : winapi.HCURSOR(winapi.LoadImage(0, winapi.MAKEINTRESOURCE(winapi.IDC_CROSS), winapi.IMAGE_CURSOR, 0, 0, winapi.LR_DEFAULTCOLOR)),
		HbrBackground : winapi.HBRUSH(winapi.GetStockObject(winapi.WHITE_BRUSH)),
		LpszMenuName  : nil,
		LpszClassName : pc,
	}
	if int(winapi.RegisterClassEx(&wc)) == 0{
		return
	}
	ws := winapi.CreateWindowEx(
		winapi.WS_EX_TOPMOST,
		pc,
		pt,
		winapi.WS_POPUP,
		20, 20,
		200, 200,
		0, 0, ins, nil,
	)
	
	winapi.ShowWindow(ws, winapi.SW_SHOW)
	winapi.SetTimer(ws, 1, 50, 0)
	
	var msg winapi.MSG
	for winapi.GetMessage(&msg, ws, 0, 0) > 0{
		winapi.TranslateMessage(&msg)
		winapi.DispatchMessage(&msg)
	}
	
	fmt.Println("\ndone.")
}

func getuintptr(str string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}
 
func windowProc(hWnd winapi.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr{
	switch msg {
	case winapi.WM_NCHITTEST:
		lr := winapi.DefWindowProc(hWnd, msg, wParam, lParam)
		if lr == HTCLIENT{
			return HTCAPTION
		}
		return lr
		
	case winapi.WM_LBUTTONDOWN:
	case winapi.WM_NCRBUTTONDOWN:
		winapi.PostMessage(hWnd, winapi.WM_CLOSE, 0, 0)
		return 0
	case winapi.WM_PAINT:
		var rc winapi.RECT 
		var ps winapi.PAINTSTRUCT 
		
		hdc := winapi.BeginPaint(hWnd, &ps)
		winapi.GetClientRect(hWnd, &rc)
		winapi.FillRect(
			hdc,
			&rc,
			(winapi.HBRUSH)(winapi.GetStockObject(winapi.GRAY_BRUSH)),
		)
		winapi.EndPaint(hWnd, &ps)
		
		return 0
	case winapi.WM_TIMER:
	case winapi.WM_DESTROY:
		winapi.PostQuitMessage(0)
		return 0
	default:
		return winapi.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}