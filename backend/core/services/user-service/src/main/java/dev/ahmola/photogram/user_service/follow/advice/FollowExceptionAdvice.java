package dev.ahmola.photogram.user_service.follow.advice;

import dev.ahmola.photogram.common_domain.error.BaseExceptionHandler;
import dev.ahmola.photogram.user_service.follow.controller.FollowController;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice(assignableTypes = {FollowController.class})
@Slf4j
public class FollowExceptionAdvice extends BaseExceptionHandler {
}
