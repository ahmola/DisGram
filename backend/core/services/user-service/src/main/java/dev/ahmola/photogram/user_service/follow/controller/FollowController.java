package dev.ahmola.photogram.user_service.follow.controller;

import dev.ahmola.photogram.user_service.follow.dto.FollowRequest;
import dev.ahmola.photogram.user_service.follow.dto.FollowResponse;
import dev.ahmola.photogram.user_service.follow.service.FollowService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Slf4j
@RequestMapping("/api/v1/follows")
@RequiredArgsConstructor
@RestController
public class FollowController {

    private final FollowService followService;

    @GetMapping(params = "followeeId")
    public ResponseEntity<List<Long>> readFollowers(@RequestParam Long followeeId){
        log.info("Read follower list of FolloweeId : {} ", followeeId);
        return ResponseEntity.ok(followService.getFollowersById(followeeId));
    }

    @GetMapping(params = "followerId")
    public ResponseEntity<List<Long>> readFollowees(@RequestParam Long followerId){
        log.info("Read followee list of FollowerId : {} ", followerId);
        return ResponseEntity.ok(followService.getFolloweesById(followerId));
    }

    @PostMapping
    public ResponseEntity<FollowResponse> createFollow(@RequestBody FollowRequest followRequest){
        log.info("Create follow, follower : {}, followee : {}",
                followRequest.followerId(), followRequest.followeeId());
        return new ResponseEntity<>(followService.createFollow(followRequest), HttpStatus.CREATED);
    }

    @DeleteMapping
    public ResponseEntity<Boolean> deleteFollow(@RequestBody FollowRequest followRequest){
        log.info("Delete follow, follower : {}, followee : {}",
                followRequest.followerId(), followRequest.followeeId());
        return ResponseEntity.ok(followService.deleteFollow(followRequest));
    }
}
