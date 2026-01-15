package dev.ahmola.photogram.user_service.user.controller;

import dev.ahmola.photogram.user_service.user.dto.UserRequest;
import dev.ahmola.photogram.user_service.user.dto.UserResponse;
import dev.ahmola.photogram.user_service.user.service.UserService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RequestMapping("/api/v1/user")
@RequiredArgsConstructor
@RestController
public class UserController {
    private final UserService userService;

    @GetMapping(params = "userId")
    public ResponseEntity<UserResponse> readUser(@RequestParam Long userId){
        log.info("Read User by UserId : {}", userId.toString());
        return ResponseEntity.ok(userService.getUserByUserId(userId));
    }

    @GetMapping(params = "username")
    public ResponseEntity<UserResponse> readUserByUsername(@RequestParam String username){
        log.info("Read User by Username : {}", username);
        return ResponseEntity.ok(userService.getUserByUsername(username));
    }

    @PostMapping
    public ResponseEntity<UserResponse> createUser(@RequestBody UserRequest userRequest){
        log.info("Create User : {}", userRequest);
        return new ResponseEntity<>(userService.createUser(userRequest), HttpStatus.CREATED);
    }

    @PutMapping
    public ResponseEntity<UserResponse> updateUser(@RequestBody UserRequest userRequest){
        log.info("Update User : {}", userRequest);
        return ResponseEntity.ok(userService.updateUser(userRequest));
    }

    @DeleteMapping
    public ResponseEntity<Boolean> deleteUser(@RequestParam Long userId){
        log.info("Delete User : {}", userId.toString());
        return ResponseEntity.ok(userService.deleteUser(userId));
    }
}