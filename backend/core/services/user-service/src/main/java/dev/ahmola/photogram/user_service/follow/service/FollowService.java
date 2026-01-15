package dev.ahmola.photogram.user_service.follow.service;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.user_service.follow.dto.FollowRequest;
import dev.ahmola.photogram.user_service.follow.dto.FollowResponse;
import dev.ahmola.photogram.user_service.follow.exception.FollowAlreadyExists;
import dev.ahmola.photogram.user_service.follow.exception.FollowNotFoundException;
import dev.ahmola.photogram.user_service.follow.model.Follow;
import dev.ahmola.photogram.user_service.follow.repository.FollowRepository;
import dev.ahmola.photogram.user_service.user.model.User;
import dev.ahmola.photogram.user_service.user.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.stream.Collectors;

@Slf4j
@RequiredArgsConstructor
@Service
@Transactional(readOnly = true)
public class FollowService {

    private final UserRepository userRepository;
    private final FollowRepository followRepository;

    public List<Long> getFollowersById(Long followeeId) {
        log.info("Start getFollowersById, followeeId : {}", followeeId);
        User followee = userRepository.getReferenceById(followeeId);

        return followRepository.findAllByFollowee(followee).stream()
                .map(follow -> follow.getFollower().getId())
                .collect(Collectors.toList());
    }

    public List<Long> getFolloweesById(Long followerId) {
        log.info("Start getFolloweesById, followerId : {]", followerId);
        User follower = userRepository.getReferenceById(followerId);

        return followRepository.findAllByFollower(follower).stream()
                .map(follow -> follow.getFollowee().getId())
                .collect(Collectors.toList());
    }

    @Transactional
    public FollowResponse createFollow(FollowRequest followRequest) {
        log.info("Start createFollow, followRequest : {}", followRequest);
        User follower = userRepository.getReferenceById(followRequest.followerId());
        User followee = userRepository.getReferenceById(followRequest.followeeId());

        // check if follow already exists
        log.info("Check if follow already exists");
        if (followRepository.existsByFollowerAndFollowee(follower, followee))
            throw new FollowAlreadyExists();

        // create follow
        log.info("Start creating...");
        try {
            followRepository.save(Follow.builder()
                    .follower(follower)
                    .followee(followee)
                    .build());
        }catch (Exception e){
            throw new RuntimeException(e.getMessage());
        }

        return FollowResponse.builder()
                .followerId(follower.getId())
                .followeeId(followee.getId())
                .build();
    }

    @Transactional
    public boolean deleteFollow(FollowRequest followRequest) {
        log.info("Check if follow exists");
        User followee = userRepository.getReferenceById(followRequest.followeeId());
        User follower = userRepository.getReferenceById(followRequest.followerId());

        if (!followRepository.existsByFollowerAndFollowee(follower, followee))
            throw new FollowNotFoundException();

        log.info("Start deleting...");
        try {
            followRepository.deleteByFollowerAndFollowee(follower, followee);
        }catch (Exception e){
            throw new RuntimeException(e.getMessage());
        }
        return true;
    }
}
