#include <windows.h>
#include <stdio.h>

const int PASS = 25358;
int main(){
	//パスワードをチェック
	int pass;
	scanf("%d", &pass);
	if(pass != PASS){
		printf("---pass Error---\n");
		return 1;
	}
	
	
	int ud;//up = 0, down = else
	while(true){
		scanf("%d", &ud);
		if(ud == 0)
			mouse_event(MOUSEEVENTF_LEFTDOWN,0,0,0,0);
		else
			mouse_event(MOUSEEVENTF_LEFTUP,0,0,0,0);
	}
}