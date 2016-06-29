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
	
	
	int x, y;//移動距離
	POINT pt;
	while(true){
		scanf("%d", &x);
		scanf("%d", &y);
		GetCursorPos(&pt);
		SetCursorPos(pt.x + x, pt.y + y);
	}
}