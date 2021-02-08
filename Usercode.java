// fq-class name com.corezoid.usercode.Usercode is mandatory
package com.corezoid.test.test;

import com.corezoid.gitcall.runner.api.IUsercode;
import java.util.Map;

public class Test implements IUsercode<Map<String, String>, Map<String, String>> {
    @java.lang.Override
    public Map<String, String> handle(Map<String, String> data) throws Exception {
        data.put("hello", "Hello java!");
        return data;
    }
}
