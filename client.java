
public class client
{
    public static void main(String[] args)
    {
	csInterface a = new csInterface();
	System.out.println(a.SignUp("huang","123456"));
	System.out.println(a.SignIn("huang","123456"));
	System.out.println(a.getMaintain());
	System.out.println(a.getContact());
	System.out.println(a.getAccount());
	String[] s = a.getAnnouncement("2018-01-01");
        System.out.println(s[0]);
	System.out.println(s[1]);
	System.out.println(s.length);
	a.close();
    
	csInterface b = new csInterface();
	System.out.println(b.SignUp("zhang","123456"));
	System.out.println(b.SignIn("zhang","aaa"));
	System.out.println(b.getMaintain());
	System.out.println(b.getContact());
	System.out.println(b.getAccount());
	String[] t = b.getAnnouncement("2018-1-31");
	System.out.println(t[0]);
	System.out.println(t[1]);
	System.out.println(t[2]);
	b.close();
	
    }
}
