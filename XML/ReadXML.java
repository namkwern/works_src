import java.util.*;
//XMLテキストのコントロール(要素の分解と読み取り)
public class ReadXML
{
	private ArrayList<Data> data = new ArrayList<Data>();
	
	//con(ファイル全体メッセージ、メインタグ)
	public ReadXML(String message, String main){
		message = ext(message, "<" + main + ">", "</" + main + ">");
		data.add(new Data(main, new ArrayList<Zokusei>(),message));
		data.addAll(getData(message));
		int count = 0;
		//条件全ての要素を分解できたら
		while(true){
			String inner = data.get(count).inner;
			if(inner != null && !inner.equals(""))
				data.addAll(getData(inner));
			if(count == data.size() - 1)break;
			count++;
		}
	}
	
	//要素を分解してdataに格納して返す
	public ArrayList<Data> getData(String message){
		ArrayList<Data> dList = new ArrayList<Data>();
		try{
			int index = 0;
			while(message.indexOf("<",index) != -1){
				String tag = ext(message, "<", ">", index);
				String[] str = tag.split(" ");
				ArrayList<Zokusei> zokusei = new ArrayList<Zokusei>();
				for(int n = 1; n < str.length; n++){
					String zoku = ext(str[n], "", "=");
					String zokuTi = ext(str[n], "\"", "\"");
					zokusei.add(new Zokusei(zoku,zokuTi));
				}
				String youso = str[0];
				String inner = ext(message, "<" + tag + ">", "</" + youso + ">");
				dList.add(new Data(youso, zokusei, inner));
				index = message.indexOf("</" + youso + ">", index) + youso.length() + 3;
			}
		}catch(IndexOutOfBoundsException e){
		}
		return dList;
	}
	
	
	//最初にはさまれた対象を取得
	public static String ext(String str, String strb, String strt){
		int indexb = str.indexOf(strb);
		int indext = str.indexOf(strt, indexb + strb.length());
		if(indexb == -1 || indext == -1)throw new IndexOutOfBoundsException("見つかりませんでした");
		return str.substring(indexb + strb.length(), indext);
		
	}
	//最初にはさまれた対象を取得(index以降のみ)
	public static String ext(String str, String strb, String strt, int index){
		int indexb = str.indexOf(strb, index);
		int indext = str.indexOf(strt, indexb + strb.length());
		if(indexb == -1 || indext == -1)throw new IndexOutOfBoundsException("見つかりませんでした");
		return str.substring(indexb + strb.length(), indext);
	}
	
	//要素名とname属性値を指定
	public String getData(String youso, String name){
		for(int n = 0; n < data.size(); n++){
			Data d = data.get(n);
			if(d.youso.equals(youso)){
				for(int m = 0; m < d.zokusei.size(); m++){
					Zokusei z = d.zokusei.get(m);
					if(z.name.equals("name")){
						if(z.value.equals(name)){
							return data.get(n).inner;
						}
					}
				}
			}
		}
		return null;
	}
	
	
	//タグ要素を全部表示
	public void dispAll(){
		System.out.println("---disp start---");
		for(int n = 0; n < data.size(); n++){
			Data d = data.get(n);
			System.out.println(d.youso + "\n" + d.inner);
			for(int m = 0; m < data.get(n).zokusei.size(); m++){
				Zokusei z = d.zokusei.get(m);
				System.out.println(z.name + ":" + z.value);
			}
			System.out.println();
		}
		System.out.println("---disp end---");
	}
	
	static class Data
	{
		public String youso;
		public ArrayList<Zokusei> zokusei;
		public String inner;//タグに挟まれたの値
		
		//con(要素名、属性リスト、タグ内)
		Data(String youso, ArrayList<Zokusei> zokusei, String inner){
			this.youso = youso;
			this.zokusei = zokusei;
			this.inner = inner;
		}
		
	}
	static class Zokusei
		{
			String name;
			String value;
			
			//con(属性名、属性値)
			Zokusei(String name, String value){
				this.name = name;
				this.value = value;
			}
		}
}