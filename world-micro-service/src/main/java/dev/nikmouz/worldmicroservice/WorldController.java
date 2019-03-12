package dev.nikmouz.worldmicroservice;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class WorldController {

  @GetMapping("/world")
  public String helloWorld() {
    return " world!";
  }
}
