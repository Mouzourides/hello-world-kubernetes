package dev.nikmouz.hellomicroservice;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class HelloController {

  private static final String HELLO = "Hello";

  @GetMapping("/hello")
  public String hello() {
    return HELLO;
  }

  @GetMapping("/helloworld")
  public String helloWorld() {
    RestTemplate restTemplate = new RestTemplate();
    String world;

    world = restTemplate.getForObject("http://localhost:8081/world", String.class);
    return HELLO + world;
  }
}
