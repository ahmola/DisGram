package dev.ahmola.photogram.user_service.controller;

import dev.ahmola.photogram.common_domain.dto.ApiResponse;
import dev.ahmola.photogram.user_service.dto.UserRequest;
import dev.ahmola.photogram.user_service.dto.UserResponse;
import dev.ahmola.photogram.user_service.services.UserService;
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

    @GetMapping
    public ResponseEntity<UserResponse> readUser(@RequestBody Long userId){
        log.info("Read User by UserId : {}", userId.toString());
        return ResponseEntity.ok(userService.getUserByUserId(userId));
    }

    @GetMapping
    public ResponseEntity<UserResponse> readUserByUsername(@RequestBody String username){
        log.info("Read User by Username : {}", username);
        return ResponseEntity.ok(userService.getUserByUsername(username));
    }

    @PostMapping
    public ResponseEntity<UserResponse> createUser(@RequestBody UserRequest userRequest){
        log.info("Create User : {}", userRequest);
        return ResponseEntity.created(userService.createUser(userRequest));
    }

    @PutMapping
    public ResponseEntity<UserResponse> updateUser(@RequestBody UserRequest userRequest){
        log.info("Update User : {}", userRequest);
        return ResponseEntity.ok(userService.updateUser(userRequest));
    }

    @DeleteMapping
    public ResponseEntity<ApiResponse> deleteUser(@RequestBody Long userId){
        log.info("Delete User : {}", userId.toString());
        return ResponseEntity.ok(userService.deleteUser(userId));
    }
}
