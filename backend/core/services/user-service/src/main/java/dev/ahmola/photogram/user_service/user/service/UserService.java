package dev.ahmola.photogram.user_service.user.service;

import dev.ahmola.photogram.user_service.user.exception.UserNotFoundException;
import dev.ahmola.photogram.user_service.user.exception.UsernameAlreadyExistsException;
import dev.ahmola.photogram.user_service.user.model.User;
import dev.ahmola.photogram.user_service.user.repository.UserRepository;
import dev.ahmola.photogram.user_service.user.dto.UserRequest;
import dev.ahmola.photogram.user_service.user.dto.UserResponse;
import dev.ahmola.photogram.user_service.user.util.UserMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@RequiredArgsConstructor
@Service
public class UserService {
    private final UserRepository userRepository;

    public UserResponse getUserByUserId(Long userId) {
        log.info("Start getUserByUserId for userId : {}", userId);
        User user = userRepository.findById(userId)
                .orElseThrow(UserNotFoundException::new);

        log.info("Success : {}", user.toString());
        return UserMapper.userToResponse(user);
    }

    public UserResponse getUserByUsername(String username) {
        log.info("Start getUserByUsername for username : {}", username);
        User user = userRepository.findByUsername(username)
                .orElseThrow(UserNotFoundException::new);

        log.info("Success : {}", user.toString());
        return UserMapper.userToResponse(user);
    }

    @Transactional
    public UserResponse createUser(UserRequest userRequest) {
        // duplication check
        log.info("Check if user already exists for username : {}", userRequest.username());
        if (userRepository.existsByUsername(userRequest.username()))
            throw new UsernameAlreadyExistsException(userRequest.username());

        // create & save
        log.info("Start createUser for userRequest : {}", userRequest);
        User user = userRepository.save(User.builder()
                        .bio(userRequest.bio())
                        .passwordHash(userRequest.passwordHash())
                        .avatarUrl(userRequest.avatarUrl())
                        .username(userRequest.username())
                .build());
        log.info("Success : {}", user.toString());

        return UserMapper.userToResponse(user);
    }

    @Transactional
    public UserResponse updateUser(UserRequest userRequest) {
        // Check if user exists
        log.info("Check if user is exists : {}", userRequest.id());
        if(!userRepository.existsById(userRequest.id()))
            throw new UserNotFoundException();

        // update
        log.info("Start updateUser for userRequest : {}", userRequest);
        User user = userRepository.save(UserMapper.requestToUser(userRequest));
        log.info("Success : {}", user.toString());

        return UserMapper.userToResponse(user);
    }

    @Transactional
    public Boolean deleteUser(Long userId) {
        // Check if user exists
        log.info("Check if user is exists : {}", userId);
        if(!userRepository.existsById(userId))
            throw new UserNotFoundException();

        // delete
        log.info("Start deleteUser for userId : {}", userId);
        userRepository.deleteById(userId);
        log.info("Success : {}", userId.toString());

        return true;
    }
}
