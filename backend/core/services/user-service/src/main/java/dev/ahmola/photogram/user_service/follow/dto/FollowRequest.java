package dev.ahmola.photogram.user_service.follow.dto;

import lombok.Builder;

@Builder
public record FollowRequest(
        Long followerId,
        Long followeeId) {}
