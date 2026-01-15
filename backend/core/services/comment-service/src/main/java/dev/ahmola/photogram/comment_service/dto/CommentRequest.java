package dev.ahmola.photogram.comment_service.dto;

import lombok.Builder;

@Builder
public record CommentRequest(
        Long id,
        Long userId,
        Long postId,
        String content
) { }