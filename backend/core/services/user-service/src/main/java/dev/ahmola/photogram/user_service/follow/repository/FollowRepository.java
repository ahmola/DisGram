package dev.ahmola.photogram.user_service.follow.repository;

import dev.ahmola.photogram.user_service.follow.model.FollowId;
import dev.ahmola.photogram.user_service.follow.model.Follow;
import dev.ahmola.photogram.user_service.user.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface FollowRepository extends JpaRepository<Follow, FollowId> {
    boolean existsByFollowerAndFollowee(User follower, User Followee);

    void deleteByFollowerAndFollowee(User follower, User followee);

    long countByFollower(User follower);

    long countByFollowee(User followee);

    // who is followed by user
    List<Follow> findAllByFollower(User follower);

    // who follows user
    List<Follow> findAllByFollowee(User followee);
}