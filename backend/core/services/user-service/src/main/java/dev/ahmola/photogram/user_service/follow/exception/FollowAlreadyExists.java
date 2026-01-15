package dev.ahmola.photogram.user_service.follow.exception;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class FollowAlreadyExists extends ApiException {
    public FollowAlreadyExists() {
      super(ErrorCode.DUPLICATE_ENTITY);
      log.error("Follow Already Exists");
    }

}
