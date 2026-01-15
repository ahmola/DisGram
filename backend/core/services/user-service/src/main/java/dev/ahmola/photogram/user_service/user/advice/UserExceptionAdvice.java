package dev.ahmola.photogram.user_service.user.advice;

import dev.ahmola.photogram.common_domain.error.BaseExceptionHandler;
import dev.ahmola.photogram.user_service.user.controller.UserController;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice(assignableTypes = {UserController.class})
@Slf4j
public class UserExceptionAdvice extends BaseExceptionHandler {
}