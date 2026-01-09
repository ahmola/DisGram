package dev.ahmola.photogram.user_service.controller;

import dev.ahmola.photogram.user_service.services.FollowService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RequiredArgsConstructor
@RestController
public class FollowController {

    private final FollowService followService;

    @GetMapping
    
}
