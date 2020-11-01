package observer

type UserService struct {
}

func (u *UserService) Register(telephone, password string) (int, error) {
	//...
	return 0, nil
}

type PromotionService struct {
}

func (p *PromotionService) IssueNewUserExperienceCash(userId int) {
	//...
}

type NotificationService struct {
}

func (n *NotificationService) SendInboxMessage(userId int, msg string) {

}

type UserController struct {
	userService  *UserController
	regObservers []RegObserver
}

func (u *UserController) Register(telephone, password string) (int, error) {
	//...
	userId, err := u.userService.Register(telephone, password)
	if err != nil {
		return -1, err
	}
	// 同步阻塞
	for _, observer := range u.regObservers {
		observer.HandleRegSuccess(userId)
	}
	return userId, nil
}

func (u *UserController) SetRegObservers(list []RegObserver) {
	u.regObservers = list
}

type RegObserver interface {
	HandleRegSuccess(userId int)
}

type RegPromotionObserver struct {
	promotionService *PromotionService
}

func (r *RegPromotionObserver) HandleRegSuccess(userId int) {
	r.promotionService.IssueNewUserExperienceCash(userId)
}

type RegNotificationObserver struct {
	notificationService *NotificationService
}

func (r *RegNotificationObserver) HandleRegSuccess(userId int) {
	r.notificationService.SendInboxMessage(userId, "Welcome...")
}
