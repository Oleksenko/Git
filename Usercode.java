// fq-class name com.corezoid.usercode.Usercode is mandatory
package test.corezoid.folder;

import com.corezoid.gitcall.runner.api.UsercodeHandler;
import java.util.Map;

public class Usercode implements UsercodeHandler<Map<String, String>, Map<String, String>> {
    @java.lang.Override
    public Map<String, String> handle(Map<String, String> data) throws Exception {
        data.put("hello", "Test java!");
        return data;
    }
}
