//comment is "SJIS"
package module;
import java.util.*;
//XMLテキストのコントロール(要素の分解と読み取り)
public class ReadXML
{
	private ArrayList<Data> data = new ArrayList<Data>();
	
	//con(ファイル全体メッセージ、メインタグ)
	public ReadXML(String message, String main){
		message = extB(message, "<" + main + ">", "</" + main + ">");
		data.add(new Data(message));
		int count = 0;
		int index = 0;
		Data d = data.get(count);
		String inner = d.inner;
		while(count < data.size()){
			if(d.inner.indexOf("<", index) == -1){
				d = data.get(count);
				inner = d.inner;
				count++;
			}
			if(inner.indexOf("<", index) == -1)break;
			String tag = extB(inner, "<", ">", index);
			String element = getElement(tag);
			//子要素の生成
			Data child = new Data(extB(inner, tag , "</" + element + ">", index));
			data.add(child);
			d.child.add(child);
			index = inner.indexOf("</" + element + ">", index) + element.length() + 3;
		}
	}	
	
	//最初にはさまれた対象を取得
	private static String ext(String str, String strb, String strt){
		return ext(str, strb, strt, 0);
	}
	//↑(index以降のみ)
	private static String ext(String str, String strb, String strt, int index){
		int indexb = str.indexOf(strb, index);
		int indext = str.indexOf(strt, indexb + strb.length());
		if(indexb == -1)throw new IndexOutOfBoundsException("見つかりませんでした:" + strb);
		if(indext == -1)throw new IndexOutOfBoundsException("見つかりませんでした:" + strt);
		return str.substring(indexb + strb.length(), indext);
	}
	
	//最初にはさまれた対象を取得、挟んだものを残す
	private static String extB(String str, String strb, String strt){
		return extB(str, strb, strt, 0);
	}
	//↑(index以降のみ)
	private static String extB(String str, String strb, String strt, int index){
		return strb + ext(str, strb, strt, index) + strt;
	}
	
	//タグから要素を取得
	private static String getElement(String tag){
		if(tag.indexOf("<") != -1){
			if(tag.indexOf(" ") != -1)
				return ext(tag, "<", " ");
			else
				return ext(tag, "<", ">");
		}else{
			if(tag.indexOf(" ") != -1)
				return ext(tag, "", " ");
			else
				return tag;
		}
	}
	
	
	//タグ要素を全部表示
	public void dispAll(){
		System.out.println("---disp start---");
		for(Data d: data)
			System.out.println(d.toString() + "\n");
		System.out.println("\n---disp end---");
	}
	
	//データクラス
	static class Data
	{
		public String element;//要素
		public ArrayList<Attribute> attribute = new ArrayList<Attribute>();//属性
		public String inner;//タグに挟まれたの値
		public ArrayList<Data> child = new ArrayList<Data>();
		
		private String tag;
		//con(要素名、属性リスト、タグ間)
		Data(String element, ArrayList<Attribute> attribute, String inner){
			setData(element, attribute, inner);
		}
		//con(タグ、タグ間)
		Data(String tag, String inner){
			setData(tag, inner);
		}
		//con(タグから終端タグ)
		Data(String all){
			setData(all);
		}
		
		//コンストラクタの受け流し
		//必ず実行される
		private void setData(String element, ArrayList<Attribute> attribute, String inner){
			this.element = element;
			this.attribute = attribute;
			this.inner = inner;
		}
		private void setData(String tag, String inner){
			ArrayList<Attribute> attribute = new ArrayList<Attribute>();
			String[] str = tag.split("[ ><]");
			for(int n = 1; n < str.length; n++)
				attribute.add(new Attribute(str[n]));
			String element = getElement(tag);
			setData(element, attribute, inner);
		}
		private void setData(String all){
			String tag = ext(all, "<", ">");
			String inner = ext(all, "<" + tag + ">", "</" + getElement(tag) + ">");
			setData(tag, inner);
		}
		
		public String toString(){
			String str = "";
			str += "[" + element + "]\n";
			for(Attribute att: attribute)
				str += att.toString() + "\n";
			if(child.size() != 0){
				for(Data chi: child)
					str += "child:<" + chi.element + ">\n";
				str += "\n";
			}
			str += inner;
			return str;
		}
	}
	
	//属性クラス
	static class Attribute
	{
		String name;
		String value;
		
		//con(属性名、属性値)
		Attribute(String name, String value){
			setAttribute(name, value);
		}
		//con(属性と属性値の塊)
		Attribute(String attribute){
			setAttribute(attribute);
		}
		
		//コンストラクタの受け流し
		private void setAttribute(String name, String value){
			this.name = name;
			this.value = value;
		}
		private void setAttribute(String attribute){
			String name = ext(attribute, "", "=");
			String value = ext(attribute, "\"", "\"");
			setAttribute(name, value);
		}
		
		public String toString(){
			return name + "=\"" + value + "\"";
		}
	}
}