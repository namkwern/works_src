package module;
import java.util.*;
import java.io.*;
import jp.hishidama.eval.*;

public class Strop
{
	//文字列中の文字の個数を返す 対象の文字列、数える文字列

	public static int count(String str, String chr)
	{
		int strlen = str.length();
		int chrlen = chr.length();
		int count = 0;
		int s = 0;
		while(count < strlen)
		{
			int index = str.indexOf(chr, s);
			if(index == -1)
				break;
			s = index + chrlen;
			count++;
		}
		return count;
	}

	//はさまれた対象を削除　対象、手前、後ろの文字列
	public static String del(String str, String strb, String strt)
	{
		int count = 0;
		int sb = 0, st = 0;
		int lenb = strb.length(), lent = strt.length(), lens = str.length();
		while(str.length() > count)
		{
			int indexb = str.indexOf(strb, sb);
			int indext = str.indexOf(strt, st);
			if(indexb == -1 || indext == -1)//見つからなかったら抜ける
				break;
			if(indexb < indext){
				str = str.substring(0, indexb + lenb) + str.substring(indext);
				sb = indexb + lenb;
			}
			
			st = indexb + lenb + lent;
			count++;
		}
		return str;
	}
	//はさまれた対象を取得ArrayList　対象、手前、後ろの文字列
	public static ArrayList<String> extAl(String str, String strb, String strt)
	{
		ArrayList<String> al = new ArrayList<String>();
		int count = 0;
		int sb = 0, st = 0;
		int lenb = strb.length(), lent = strt.length(), lens = str.length();
		while(str.length() > count)
		{
			int indexb = str.indexOf(strb, sb);
			int indext = str.indexOf(strt, st);
			if(indexb == -1 || indext == -1)//見つからなかったら抜ける
				break;
			if(indexb < indext){
				al.add(str.substring(indexb + lenb, indext));
				sb = indexb + lenb;
			}
			
			st = indexb + lenb;
			count++;
		}
		return al;
	}
	//最初にはさまれた対象を取得
	public static String ext(String str, String strb, String strt){
		int indexb = str.indexOf(strb);
		int indext = str.indexOf(strt, indexb);
		str = str.substring(indexb + strb.length(), indext);
		return str;
	}
	//最初にはさまれた対象を取得(index以降のみ)
	public static String ext(String str, String strb, String strt, int index){
		int indexb = str.indexOf(strb, index);
		int indext = str.indexOf(strt, indexb);
		str = str.substring(indexb + strb.length(), indext);
		return str;
	}
	//はさまれた対象を取得String　対象、手前、後ろの文字列
	public static String extStr(String str, String strb, String strt)
	{
		String rv = "";
		if(!(strb.equals(strt))){
			int count = 0;
			int sb = 0, st = 0;
			int lenb = strb.length(), lent = strt.length(), lens = str.length();
			while(str.length() > count)
			{
				int indexb = str.indexOf(strb, sb);
				int indext = str.indexOf(strt, st);
				if(indexb == -1 || indext == -1)//見つからなかったら抜ける
					break;
				if(indexb < indext){
					rv += str.substring(indexb + lenb, indext);
					sb = indexb + lenb;
				}
				
				st = indexb + lenb + lent;
				count++;
			}
		}else{
			int beg= 0;//始点
			int count = 0;
			int s = 0;
			while(str.length() > count)
			{
				int index = str.indexOf(strb, s);
				if(index == -1)//見つからなかったら抜ける
					break;
				//奇数個と偶数個の"を判断して処理
				if(count % 2 == 0){
					beg = index;
				}else{
					rv += str.substring(beg + 1, index);
				}
				s = str.indexOf(strb, s) + 1;
				count++;
			}
		}
		return rv;
	}
	//はさまれた(前後文字列を含む)対象を取得String　対象、手前、後ろの文字
	public static String extStrp(String str, String strb, String strt, int eflag)
	{
		
		String rv = "";
		int count = 0;
		int s = 0;
		int lenb = strb.length(), lent = strt.length(), lens = str.length();
		int indexb = 0, indext = 0;
		while(str.length() > count)
		{
			
			//エスケープするか？
			if(eflag == 1){
				indexb = str.indexOf(strb, s);
				
				while(indexb > 0 && str.substring(indexb - 1).indexOf("\\") == 0){
					
					indexb = str.indexOf(strb, indexb + 1);
				}
				indext = str.indexOf(strt, indexb + 1);
				while(indext > 0 && str.substring(indext - 1).indexOf("\\") == 0){
					indext = str.indexOf(strt, indext + 1);
				}
			}else{
				indexb = str.indexOf(strb, s);
				indext = str.indexOf(strt, indexb + 1);
			}
			if(indexb == -1 || indext == -1)//見つからなかったら抜ける
				break;
			//抽出
			if(indexb < indext){
				String a = str.substring(indexb + lenb, indext);
				//空白をエスケープ
				if(a.length() == 0)
					a = "escspacef903q0gjaw2r09";
				rv += strb + a + strt + "\r\n";
				s = indext + lenb;
			}
			
			count++;
		}
		
		return rv;
	}
		//置換の代替 対象の文字列、置換前、置換後
	public static String priRep(String str, String bef, String aft){
		//System.out.println(str + " " + bef + " " + aft);
		int count = 0;
		int s = 0;
		int lenb = bef.length(), lena = aft.length(), lens = str.length();
		while(lens > count)
		{
			int indexb = str.indexOf(bef, s);
			if(indexb == -1)//見つからなかったら抜ける
				break;
			str = str.substring(0, indexb) + aft + str.substring(indexb + lenb);
			
			s = indexb + lena;
			count++;
		}
		
		return str;
	}
	//置換の代替 対象の文字列、置換前、置換後
	public static String priRepfor(String str, String bef, String aft){
		String aft2 = "";
		int repcount = 0;
		int count = 0;
		int s = 0;
		int lens = str.length();
		while(lens > count)
		{
			String bef2 = forn(bef, "#for(", ")#");
			int indexb = str.indexOf(bef2, s);
			if(indexb == -1)//見つからなかったら抜ける
				break;
			
			//for文解釈
			fornAdd();
			aft2 = forn2(aft, "#for(", ")#");
			str = str.substring(0, indexb) + aft2 + str.substring(indexb + bef2.length());
			s = indexb + aft2.length();
			
			count++;
		}
		
		return str;
	}
	//文章置換処理　対象、置換前、置換後、ワイルドカード
	public static String repWild(String str, String bef, String aft, String wild, int eflag)
	{
		if(aft.indexOf(wild) == -1){//befワイルドカードあり
			int count = 0;
			int s = 0;
			
			int indexb = 0, indext = 0;
			while(str.length() > count)
			{
				String bef2 = forn(bef, "#for(", ")#");
				
				String strb = beg(bef2, wild);
				String strt = tar(bef2, wild);
				int lenb = strb.length(), lent = strt.length();
				//エスケープするかな
				if(eflag == 1){
					
					indexb = str.indexOf(strb, s);
					while(indexb > 0 && str.substring(indexb - 1).indexOf("\\") == 0){
						indexb = str.indexOf(strb, indexb + 1);
					}
					indext = str.indexOf(strt, indexb + 1);
					while(indext > 0 && str.substring(indext - 1).indexOf("\\") == 0){
						indext = str.indexOf(strt, indext + 1);
					}
				}else{
					indexb = str.indexOf(strb, s);
					indext = str.indexOf(strt, indexb + 1);
				}
				//前か後ろが未指定の場合に対処する
				if(strb.length() == 0)
					indexb = 0;
				if(strt.length() == 0)
					indext = str.length();
				if(indexb == -1 || indext == -1)break;//見つからなかったら抜ける
				
				//for文解釈
				fornAdd();
				String aft2 = forn2(aft, "#for(", ")#");
				str = str.substring(0, indexb) + aft2 + str.substring(indext + lent);
				s = indexb + aft2.length();
					
				
				count++;
				if(strb.length() == 0 && strt.length() == 0)break;
			}
			
		}else{//bef,aftにワイルドカードあり
			int count = 0;
			int s = 0;
			while(str.length() > count)
			{
				String bef2 = forn(bef, "#for(", ")#");
				
				String strb = beg(bef2, wild);
				String strt = tar(bef2, wild);
				int lenb = strb.length(), lent = strt.length();
				//エスケープするかな
				int indexb = 0, indext = 0;
				if(eflag == 1){
					indexb = str.indexOf(strb, s);
					while(indexb > 0 && str.substring(indexb - 1).indexOf("\\") == 0){
						
						indexb = str.indexOf(strb, indexb + 1);
					}
					indext = str.indexOf(strt, indexb + 1);
					while(indext > 0 && str.substring(indext - 1).indexOf("\\") == 0){
						indext = str.indexOf(strt, indext + 1);
					}
				}else{
					indexb = str.indexOf(strb, s);
					indext = str.indexOf(strt, indexb + 1);
				}
				//前か後ろが未指定の場合に対処する
				if(strb.length() == 0){
					indexb = 0;
				}
				if(strt.length() == 0){
					indext = str.length();
				}
				
				//見つからなかったら抜ける
				if(indexb == -1 || indext == -1)break;
				
				String onestr = str.substring(indexb + lenb, indext);
				String str1 = str.substring(0, indexb);
				str = str.substring(indext + lent);
				
				//for文解釈
				fornAdd();
				String aft2 = forn2(aft, "#for(", ")#");
				str = str1 + priRep(aft2, wild, onestr) + str;
				s = indexb + priRep(aft2, wild, onestr).length();
				
				count++;
				if(strb.length() == 0 && strt.length() == 0)break;
			}
			
		}
		return str;
	}
	//文字列の抽出
	public static String extWild(String str, String bef, String wild, int eflag)
	{
		int count = 0;
		int s = 0;
		int indexb = 0, indext = 0;
		String rv = "";
		while(str.length() > count)
		{
			String bef2 = forn(bef, "#for(", ")#");
			
			String strb = beg(bef2, wild);
			String strt = tar(bef2, wild);
			int lenb = strb.length(), lent = strt.length();
			//エスケープするかな
			if(eflag == 1){
				
				indexb = str.indexOf(strb, s);
				while(indexb > 0 && str.substring(indexb - 1).indexOf("\\") == 0){
					indexb = str.indexOf(strb, indexb + 1);
				}
				indext = str.indexOf(strt, indexb + 1);
				while(indext > 0 && str.substring(indext - 1).indexOf("\\") == 0){
					indext = str.indexOf(strt, indext + 1);
				}
			}else{
				indexb = str.indexOf(strb, s);
				indext = str.indexOf(strt, indexb + 1);
			}
			//前か後ろが未指定の場合に対処する
			if(strb.length() == 0)
				indexb = 0;
			if(strt.length() == 0)
				indext = str.length();
			if(indexb == -1 || indext == -1)break;//見つからなかったら抜ける
			//for文解釈
			fornAdd();
			rv += strb + str.substring(indexb + lenb, indext) + strt;
			s = indexb + lenb;
			count++;
			if(strb.length() == 0 && strt.length() == 0)break;
		}
		return rv;
	}
	//#for()#の中身を抜き出して計算して置き換える
	private static int repcount = 0;
	public static void resetrep(){repcount = 0;repcount2 = 0;}
	public static String forn(String aft, String beg, String tar){
		if(aft.indexOf("#for(") != -1 && aft.indexOf(")#") != -1){
			String formu = extStr(aft, beg, tar);
			formu = priRep(formu, "n", String.valueOf(repcount));//数式を取り出す
			Rule rule = ExpRuleFactory.getDefaultRule();//インスタンス化
			Expression exp = rule.parse(formu);	//解析
			long result = exp.evalLong(); 	//計算実施
			String wc = extStrp(aft, beg, tar, 0);
			wc = priRep(wc, "\r\n", "");
			return priRep(aft, wc, String.valueOf(result));
		}else
			return aft;
	}
	public static void fornAdd(){repcount++;}
	private static int repcount2 = 0;
	public static String forn2(String aft, String beg, String tar){
		if(aft.indexOf("#for(") != -1 && aft.indexOf(")#") != -1){
			String formu = extStr(aft, beg, tar);
			formu = priRep(formu, "n", String.valueOf(repcount2++));//数式を取り出す
			Rule rule = ExpRuleFactory.getDefaultRule();//インスタンス化
			Expression exp = rule.parse(formu);	//解析
			long result = exp.evalLong(); 	//計算実施
			String wc = extStrp(aft, beg, tar, 0);
			wc = priRep(wc, "\r\n", "");
			return priRep(aft, wc, String.valueOf(result));
		}else
			return aft;
	}
	public static String beg(String str, String chr){
		if(str.indexOf(chr) != -1)
			str = str.substring(0, str.indexOf(chr));
		else
			str = "";
		return str;
	}
	public static String tar(String str, String chr){
		if(str.indexOf(chr) != -1)
			str = str.substring(str.indexOf(chr) + chr.length());
		else
			str = "";
		return str;
	}
	//改行単位で分割する(空白行カット処理)　分割する文字列
	public static String[] splitln(String str){
		str += "\n";
		ArrayList<String> al2 = new ArrayList<String>();
		int index = 0;
		int h = -1;
		int cou = 0;
		while(true){
			index = str.indexOf("\n", h + 1);
			if(index == -1)break;
				String a = str.substring(h + 1, index);
				if(a.length() != 0)
					al2.add(a);
			h = index;
			cou++;
		}
		String[] rv = al2.toArray(new String[al2.size()]);
		return rv;
	}
	//String型を次の値に進める(数字→大文字→小文字)
	public static String nextString(String str){
		int leng = str.length() - 1;
		char[] c = str.toCharArray();//Stringを複写
		int n;
		for(n = 0 ; n < leng + 1; n++){
			int index = leng - n;
			if(c[index] == '9')
				c[index] += 'A' - '9';		//9をAに変換
			else if(c[index] == 'Z')
				c[index] += 'a' - 'Z';		//Zをaに変換
			else if(c[index] == 'z'){
				c[index] += '0' - 'z';		//zを0に変換
				continue;			//次の文字に進める(桁上がり)
			}
			else
				c[index]++;			//文字を一つ進める
			break;					//continueが無ければ終了
		}
		str = String.valueOf(c);
		if(n - 1 == leng && c[0] == '0')str = '0' + str;	//最大桁で桁上がりしたら先頭に0追加
		return str;
	}
	
	//一般的なコマンドライン引数の妥当性を確認するための正規表現を返す
	//"ops"→-op -p -pos などを正しいものとする
	public String getSeiki(String ops){
		String seiki = "-([" + ops + "])(?!\\1)";
		for(int n = 2; n < ops.length(); n++)
		{
			seiki += "([" + ops + "])?";
			for(int m = 1; m <= n; m++)
				seiki += "(?!\\" + m + ")";
		}
		seiki += "[" + ops + "]?$";
		return seiki;
	}
}
